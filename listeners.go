package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lcallbacks -llisteners
// #include "callbacks.h"
// #include "listeners.h"
import "C"

import "unsafe"

// Registry Listener
//
//export registryHandleGlobalCallback
func registryHandleGlobalCallback(data unsafe.Pointer, registry *C.struct_wl_registry, name C.uint32_t, interfaceName *C.char, version C.uint32_t) {
	// Handle registry global event
}

//export registryHandleGlobalRemoveCallback
func registryHandleGlobalRemoveCallback(data unsafe.Pointer, registry *C.struct_wl_registry, name C.uint32_t) {
	// Handle registry global remove event
}

var registry_listener = C.struct_wl_registry_listener{
	global:        (*[0]byte)(C.registryHandleGlobal),
	global_remove: (*[0]byte)(C.registryHandleGlobalRemove),
}

// Shell Surface Listener
//
//export shellSurfaceHandlePingCallback
func shellSurfaceHandlePingCallback(data unsafe.Pointer, shellSurface *C.struct_wl_shell_surface, serial C.uint32_t) {
	// Handle shell surface ping event
}

//export shellSurfaceHandleConfigureCallback
func shellSurfaceHandleConfigureCallback(data unsafe.Pointer, shellSurface *C.struct_wl_shell_surface, edges C.uint32_t, width C.int32_t, height C.int32_t) {
	// Handle shell surface configure event
}

//export shellSurfaceHandlePopupDoneCallback
func shellSurfaceHandlePopupDoneCallback(data unsafe.Pointer, shellSurface *C.struct_wl_shell_surface) {
	// Handle shell surface popup done event
}

func initRegistry(tinywl *tinywl) {
	C.wl_registry_add_listener(tinywl.registry, &C.registry_listener, unsafe.Pointer(tinywl))
}

func initShellSurface(tinywl *tinywl, surface *C.struct_wl_surface) {
	tinywl.shellSurface = C.wl_shell_get_shell_surface(tinywl.shell, surface)
	C.wl_shell_surface_add_listener(tinywl.shellSurface, &C.shell_surface_listener, unsafe.Pointer(tinywl))
}
