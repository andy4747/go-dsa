# File: CMakeLists.txt
cmake_minimum_required(VERSION 3.10)
project(dsacpp)

set(CMAKE_CXX_STANDARD 11)

# Add executables
add_executable(main main.cpp data_structures/array.cpp)
