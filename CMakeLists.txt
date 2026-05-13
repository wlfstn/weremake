cmake_minimum_required(VERSION 4.1.0)
project(weremake VERSION 0.1.0)

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_CXX_SCAN_FOR_MODULES ON)
set(CMAKE_CXX_FLAGS_RELEASE "-O3 -DNDEBUG")
set(CMAKE_CXX_FLAGS_DEBUG "-g -O0")

add_compile_options(-Wno-c23-extensions)

set(APP_ICO "${CMAKE_SOURCE_DIR}/src/embed/resource/icon.rc")

set(SOURCE_FILES 
	src/main.cpp
)
set_source_files_properties(${SOURCE_FILES} PROPERTIES
	COMPILE_OPTIONS "$<$<CONFIG:Release>:-fno-exceptions>"
)

set(MODULE_FILES
	src/core/cli.cppm
)
set_source_files_properties(${MODULE_FILES} PROPERTIES
	COMPILE_OPTIONS "$<$<CONFIG:Release>:-fno-exceptions>"
)

# set(WERELIB_SHARED ON CACHE BOOL "" FORCE)
add_subdirectory(src/vendor/werelib)

add_executable(${PROJECT_NAME} ${SOURCE_FILES})
target_include_directories(${PROJECT_NAME} PRIVATE src/include src/vendor/stb)
target_compile_definitions(${PROJECT_NAME} PRIVATE
	EMBED_ROOT=${CMAKE_CURRENT_SOURCE_DIR}/src/embed
	UNICODE
	_UNICODE
)

target_sources(${PROJECT_NAME} PRIVATE
	FILE_SET cxx_modules TYPE CXX_MODULES FILES
		${MODULE_FILES}
)

target_link_libraries(
	${PROJECT_NAME} PRIVATE
	werelib
)