cmake_minimum_required(VERSION 3.0.0)
project(graphite VERSION 0.1.0)

include(CTest)
enable_testing()

# LLVM - PART A
find_package(LLVM REQUIRED CONFIG)
include_directories(${LLVM_INCLUDE_DIRS})
add_definitions(${LLVM_DEFINITIONS})
# End of LLVM - PART A

add_library(graphite graphite.cpp)

# LLVM - PART B
llvm_map_components_to_libnames(llvm_libs support core irreader)
target_link_libraries(graphite ${llvm_libs})
# End of LLVM - PART B