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

#endif //BUGU_OBFUSION_SERVICE_INCLUDE_BUGU_OBFUSION_BUGU_OBFUSION_H_
