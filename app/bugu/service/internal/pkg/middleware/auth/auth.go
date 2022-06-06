/*
 * MIT License
 *
 * Copyright (c) 2022. HominSu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package auth

import (
	"context"
	nethttp "net/http"
	"strings"
	"time"

	pkg "github.com/hominsu/bugu/app/bugu/service/internal/pkg/http/error"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

const (

	// bearerWord the bearer key word for authorization
	bearerWord string = "Token"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

func GenerateToken(secret, userid string) (string, error) {
	expireTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))
	claims := &Claims{
		UserID: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "bugu.cn",
			Subject:   "bugu user token",
			ExpiresAt: expireTime,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func JwtAuthServiceMiddleware(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				jwtToken, err := getJwtTokenFromTrans(header)
				if err != nil {
					return nil, err
				}

				claims := &Claims{}
				tokenInfo, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, ErrUnSupportSigningMethod
					}
					return []byte(secret), nil
				})
				if err != nil {
					if ve, ok := err.(*jwt.ValidationError); ok {
						if ve.Errors&jwt.ValidationErrorMalformed != 0 {
							return nil, ErrTokenInvalid
						} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
							return nil, ErrTokenExpired
						} else {
							return nil, ErrTokenParseFail
						}
					}
					return nil, errors.Unauthorized(reason, err.Error())
				} else if !tokenInfo.Valid {
					return nil, ErrTokenInvalid
				}

				// append the userid to ctx for next service
				ctx = metadata.AppendToClientContext(ctx, "x-md-global-userid", claims.UserID)
			}
			return handler(ctx, req)
		}
	}
}

func getJwtTokenFromTrans(tr transport.Transporter) (string, error) {
	auths := strings.SplitN(tr.RequestHeader().Get(authorizationKey), " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
		return "", ErrMissingJwtToken
	}
	return auths[1], nil
}

func JwtAuthRouteFilter(secret string) func(nethttp.Handler) nethttp.Handler {
	return func(next nethttp.Handler) nethttp.Handler {
		return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
			jwtToken, err := getJwtTokenFromReq(r)
			if err != nil {
				pkg.ErrorEncoder(w, r, &pkg.HTTPError{Code: 401, Message: "JWT token is missing"})
				return
			}

			claims := &Claims{}
			tokenInfo, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, ErrUnSupportSigningMethod
				}
				return []byte(secret), nil
			})
			if err != nil {
				if ve, ok := err.(*jwt.ValidationError); ok {
					if ve.Errors&jwt.ValidationErrorMalformed != 0 {
						pkg.ErrorEncoder(w, r, &pkg.HTTPError{Code: 401, Message: "Token is invalid"})
						return
					} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						pkg.ErrorEncoder(w, r, &pkg.HTTPError{Code: 401, Message: "JWT token has expired"})
						return
					} else {
						pkg.ErrorEncoder(w, r, &pkg.HTTPError{Code: 401, Message: "Fail to parse JWT token"})
						return
					}
				}
				pkg.ErrorEncoder(w, r, &pkg.HTTPError{Code: 401, Message: err.Error()})
				return
			} else if !tokenInfo.Valid {
				pkg.ErrorEncoder(w, r, &pkg.HTTPError{Code: 401, Message: "Token is invalid"})
			}

			// append the userid to ctx for next service
			r.Header.Set("x-md-global-userid", claims.UserID)

			next.ServeHTTP(w, r)
		})
	}
}

func getJwtTokenFromReq(r *nethttp.Request) (string, error) {
	auths := strings.SplitN(r.Header.Get(authorizationKey), " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
		return "", ErrMissingJwtToken
	}
	return auths[1], nil
}
