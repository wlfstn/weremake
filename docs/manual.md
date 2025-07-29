## Manual on WereMake

Start a new project with `weremake init projectName`

## TOML CONFIG
```toml
PROJECT_NAME = "DmxRasterizer"
CXX_STANDARD = 20

SOURCE = [
	"main.cpp",
	"shader.cpp"
]
HEADER = [
	"{VAR}/GLAD/include",
	"{VAR}/GLFW/include",
	"{VAR}/GLM"
]

[CREATE_STATIC]
glad = "${VAR}/glad/src/glad.c"

[LINK]
STATIC = [
"glad",
"{VAR}/GLFW/lib/glfw3.lib"
]
```
