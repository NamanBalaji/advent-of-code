cmake_minimum_required(VERSION 3.10)
project(AdventOfCode)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Include utils
include_directories(utils)

# Add executables for each day
add_executable(Day01 Day01/main.cpp utils/file_utils.cpp)
add_executable(Day02 Day02/main.cpp utils/file_utils.cpp)
add_executable(Day03 Day03/main.cpp utils/file_utils.cpp)
add_executable(Day04 Day04/main.cpp utils/file_utils.cpp)
add_executable(Day05 Day05/main.cpp utils/file_utils.cpp)