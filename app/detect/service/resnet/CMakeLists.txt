cmake_minimum_required(VERSION 3.11)

project(resnet
        LANGUAGES C CXX)
option(CUDA_USE_STATIC_CUDA_RUNTIME OFF)
set(CMAKE_CXX_STANDARD 17)

# Cuda
include_directories(/usr/local/cuda/include)
link_directories(/usr/local/cuda/lib64)

# Tensorrt
include_directories(/usr/include/x86_64-linux-gnu/)
link_directories(/usr/lib/x86_64-linux-gnu/)

if (CMAKE_BUILD_TYPE STREQUAL Debug)
    ADD_DEFINITIONS(-DDEBUG)
    message(STATUS "CMake Build Type: Debug")
    message("")
elseif (CMAKE_BUILD_TYPE STREQUAL Release)
    ADD_DEFINITIONS(-DRELEASE)
    message(STATUS "CMake Build Type: Release")
    message("")
endif ()

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Werror -Wno-unused-parameter")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O2")

# Threads
find_package(Threads REQUIRED)
if (Threads_FOUND)
    message("")
endif ()

include_directories(../include)
include_directories(include)
file(GLOB ProjectSRC
        "src/*.cc")

add_executable(${PROJECT_NAME} ${ProjectSRC})
target_link_libraries(${PROJECT_NAME}
        Threads::Threads
        nvinfer
        cudart)
