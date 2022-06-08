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

package pkg

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (ss *Sha1Stream) UpdateByte(data []byte) (size int64, err error) {
	if ss._sha1 == nil {
		ss._sha1 = sha1.New()
	}

	rSize, err := ss._sha1.Write(data)
	if err != nil {
		return 0, err
	}

	return int64(rSize), nil
}

func (ss *Sha1Stream) UpdateIO(ioReader io.Reader) (size int64, err error) {
	if ss._sha1 == nil {
		ss._sha1 = sha1.New()
	}

	rSize, err := io.Copy(ss._sha1, ioReader)
	if err != nil {
		return 0, err
	}

	return rSize, nil
}

func (ss *Sha1Stream) Reset() {
	ss._sha1.Reset()
}

func (ss *Sha1Stream) Sum() string {
	return hex.EncodeToString(ss._sha1.Sum([]byte("")))
}

func Sha1(data []byte) (string, int64, error) {
	_sha1 := sha1.New()
	size, err := _sha1.Write(data)
	if err != nil {
		return "", 0, err
	}
	return hex.EncodeToString(_sha1.Sum([]byte(""))), int64(size), nil
}

func IOSha1(ioReader io.Reader) (string, int64, error) {
	_sha1 := sha1.New()
	size, err := io.Copy(_sha1, ioReader)
	if err != nil {
		return "", 0, err
	}
	return hex.EncodeToString(_sha1.Sum([]byte(""))), size, nil
}

func MD5(data []byte) (string, int64, error) {
	_md5 := md5.New()
	size, err := _md5.Write(data)
	if err != nil {
		return "", 0, err
	}
	return hex.EncodeToString(_md5.Sum([]byte(""))), int64(size), nil
}

func IOMd5(file io.Reader) (string, int64, error) {
	_md5 := md5.New()
	size, err := io.Copy(_md5, file)
	if err != nil {
		return "", 0, err
	}
	return hex.EncodeToString(_md5.Sum(nil)), size, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSize(filename string) (int64, error) {
	var result int64
	err := filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}
