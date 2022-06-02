// MIT License
//
// Copyright (c) 2022. HominSu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

//
// Created by Homin Su on 2022/5/17.
//

#ifndef BUGU_OBFUSION_SERVICE_INCLUDE_BUGU_OBFUSION_BUGU_OBFUSION_H_
#define BUGU_OBFUSION_SERVICE_INCLUDE_BUGU_OBFUSION_BUGU_OBFUSION_H_

#if defined(__has_builtin)
#define BUGU_HAS_BUILTIN(x) __has_builtin(x)
#else
#define BUGU_HAS_BUILTIN(x) 0
#endif

#ifndef BUGU_ASSERT
#include <cassert>
#define BUGU_ASSERT(x) assert(x)
#endif // BUGU_ASSERT

#if defined(_WIN64) || defined(WIN64) || defined(_WIN32) || defined(WIN32)
#if defined(_WIN64) || defined(WIN64)
#define BUGU_ARCH_64 1
#else
#define BUGU_ARCH_32 1
#endif
#define BUGU_PLATFORM_STRING "windows"
#define BUGU_WINDOWS 1
#elif defined(__linux__)
#define BUGU_PLATFORM_STRING "linux"
#define BUGU_LINUX 1
#ifdef _LP64
#define BUGU_ARCH_64 1
#else /* _LP64 */
#define BUGU_ARCH_32 1
#endif /* _LP64 */
#elif defined(__APPLE__)
#define BUGU_PLATFORM_STRING "osx"
#define BUGU_APPLE 1
#ifdef _LP64
#define BUGU_ARCH_64 1
#else /* _LP64 */
#define BUGU_ARCH_32 1
#endif /* _LP64 */
#endif

#ifndef BUGU_WINDOWS
#define BUGU_WINDOWS 0
#endif
#ifndef BUGU_LINUX
#define BUGU_LINUX 0
#endif
#ifndef BUGU_APPLE
#define BUGU_APPLE 0
#endif

#ifdef _MSC_VER
#if _MSC_VER < 1700
typedef __int8 int8_t;
typedef __int16 int16_t;
typedef __int32 int32_t;
typedef __int64 int64_t;
typedef unsigned __int8 uint8_t;
typedef unsigned __int16 uint16_t;
typedef unsigned __int32 uint32_t;
typedef unsigned __int64 uint64_t;
#else
#include <stdint.h>
#endif /* _MSC_VER < 1700 */
#else
#include <stdint.h>
#endif /* _MSC_VER */

#ifdef _MSC_VER
#if _MSC_VER < 1400
#define PRIdS __PRIS_PREFIX "d"
#define PRIxS __PRIS_PREFIX "x"
#define PRIuS __PRIS_PREFIX "u"
#define PRIXS __PRIS_PREFIX "X"
#define PRIoS __PRIS_PREFIX "o"
#else
#include <inttypes.h>
#endif /* _MSC_VER < 1400 */
#else
#include <inttypes.h>
#endif /* _MSC_VER */

#endif //BUGU_OBFUSION_SERVICE_INCLUDE_BUGU_OBFUSION_BUGU_OBFUSION_H_
