# Threads
find_package(Threads REQUIRED)
if (Threads_FOUND)
    message("")
endif ()

# OpenSSL
set(OPENSSL_USE_STATIC_LIBS true)
find_package(OpenSSL REQUIRED)
if (OPENSSL_FOUND)
    message(STATUS "OPENSSL_VERSION = ${OPENSSL_VERSION}")
    message(STATUS "OPENSSL_INCLUDE_DIR = ${OPENSSL_INCLUDE_DIR}")
    message(STATUS "OPENSSL_LIBRARIES = ${OPENSSL_LIBRARIES}")
endif ()

# Protobuf
find_package(Protobuf REQUIRED)
if (Protobuf_FOUND)
    message(STATUS "Protobuf_VERSION = ${Protobuf_VERSION}")
    message(STATUS "Protobuf_INCLUDE_DIRS = ${Protobuf_INCLUDE_DIRS}")
    message(STATUS "Protobuf_LIBRARY = ${Protobuf_LIBRARY}")
endif ()

# gRPC
find_package(gRPC CONFIG REQUIRED)
if (gRPC_FOUND)
    message(STATUS "gRPC_VERSION = ${gRPC_VERSION}")
endif ()

include_directories(${PROJECT_SOURCE_DIR}/include)
include_directories(${PROJECT_SOURCE_DIR}/third_party)
include_directories(${PROJECT_SOURCE_DIR}/../../../)
file(GLOB ProjectSRC
        "../../../api/obfusion/service/v1/cpp/*.cc"
        "thread_pool/*.cc"
        "thread/*.cc"
        "*.cc")

if (CMAKE_BUILD_TYPE STREQUAL Debug)
    ADD_DEFINITIONS(-DDEBUG)
    message(STATUS "CMake Build Type: Debug")
    message("")
elseif (CMAKE_BUILD_TYPE STREQUAL Release)
    ADD_DEFINITIONS(-DRELEASE)
    message(STATUS "CMake Build Type: Release")
    message("")
endif ()

add_executable(${PROJECT_NAME} ${ProjectSRC})
target_link_libraries(${PROJECT_NAME}
        Threads::Threads
        protobuf
        gRPC::grpc++
        obfusion-static)
