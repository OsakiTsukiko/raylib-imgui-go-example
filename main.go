package main

/*
#cgo CFLAGS: -I ./rlImGui -I ./cimgui
#cgo LDFLAGS: -L ./rlImGui -L ./cimgui -lrlImGui -lcimgui -lstdc++
#include "rlImGui.h"

extern void igShowDemoWindow(bool* p_open);
*/
import "C"
import (
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1080, 720, "Raylib ImGui")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	C.rlImGuiSetup(true)
	defer C.rlImGuiShutdown()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Red)

		C.rlImGuiBegin()

		// Example: Use ImGui functions directly
		// C.TextUnformatted(C.CString("Hello from rlImGui!"))

		copen := C.bool(false)
		copen_ptr := (*C.bool)(unsafe.Pointer(&copen))
		C.igShowDemoWindow(copen_ptr)

		// End ImGui frame
		C.rlImGuiEnd()
		// rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
