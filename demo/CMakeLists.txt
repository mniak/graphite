cmake_minimum_required(VERSION 3.0.0)
project(graphite VERSION 0.1.0)

include(CTest)
enable_testing()

add_executable(graphite-demo demo.cpp)

target_link_libraries (graphite-demo LINK_PUBLIC graphite)