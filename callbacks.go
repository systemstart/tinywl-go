package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lcallbacks -llisteners
// #cgo !windows CFLAGS: -DXKB_COMPOSE -DXKB_TEXT -DXKB_ATOMS -DXKB_VMOD -DXKB_MODIFIERS -DXKB_GLIB -DXKB_KEYMAP_FLAGS
// #cgo !windows LDFLAGS: -lxkbcommon
// #include "callbacks.h"
// #include "listeners.h"
// #ifndef _WIN32
//   #include <xkbcommon/xkbcommon.h>
// #endif
import "C"

import "C"
import (
	"log"
	"unsafe"
)

type seatListenerData struct {
	tinywl *tinywl
}

type touchListenerData struct {
	tinywl *tinywl
}

type keyboardListenerData struct {
	tinywl *tinywl
}

// Seat Listener
//
//export seatHandleCapabilitiesCallback
func seatHandleCapabilitiesCallback(data unsafe.Pointer, seat *C.struct_wl_seat, capabilities C.uint32_t) {
	listenerData := (*seatListenerData)(data)
	// Handle seat capabilities event
	log.Printf("listenerData: %#v", listenerData)
}

//export seatHandleNameCallback
func seatHandleNameCallback(data unsafe.Pointer, seat *C.struct_wl_seat, name *C.char) {
	listenerData := (*seatListenerData)(data)
	// Handle seat name event
	log.Printf("listenerData: %#v", listenerData)
}

// Pointer Listener
//
//export pointerHandleEnterCallback
func pointerHandleEnterCallback(data unsafe.Pointer, pointer *C.struct_wl_pointer, serial C.uint32_t, surface *C.struct_wl_surface, surface_x C.wl_fixed_t, surface_y C.wl_fixed_t) {
	listenerData := (*pointerListenerData)(data)
	// Handle pointer enter event
	log.Printf("listenerData: %#v", listenerData)
}

//export pointerHandleLeaveCallback
func pointerHandleLeaveCallback(data unsafe.Pointer, pointer *C.struct_wl_pointer, serial C.uint32_t, surface *C.struct_wl_surface) {
	listenerData := (*pointerListenerData)(data)
	// Handle pointer leave event
	log.Printf("listenerData: %#v", listenerData)
}

// Keyboard Listener
//
//export keyboardHandleKeymapCallback
func keyboardHandleKeymapCallback(data unsafe.Pointer, keyboard *C.struct_wl_keyboard, format C.uint32_t, fd C.int32_t, size C.uint32_t) {
	listenerData := (*keyboardListenerData)(data)
	// Handle keyboard keymap event
	log.Printf("listenerData: %#v", listenerData)
}

//export keyboardHandleEnterCallback
func keyboardHandleEnterCallback(data unsafe.Pointer, keyboard *C.struct_wl_keyboard, serial C.uint32_t, surface *C.struct_wl_surface, keys *C.struct_wl_array) {
	listenerData := (*keyboardListenerData)(data)
	// Handle keyboard enter event
	log.Printf("listenerData: %#v", listenerData)
}

// Touch Listener
//
//export touchHandleDownCallback
func touchHandleDownCallback(data unsafe.Pointer, touch *C.struct_wl_touch, serial C.uint32_t, time C.uint32_t, surface *C.struct_wl_surface, id C.int32_t, x C.wl_fixed_t, y C.wl_fixed_t) {
	listenerData := (*touchListenerData)(data)
	// Handle touch down event
	log.Printf("keyboardHandleKeymap: %#v", listenerData)
}

//export touchHandleUpCallback
func touchHandleUpCallback(data unsafe.Pointer, touch *C.struct_wl_touch, serial C.uint32_t, time C.uint32_t, id C.int32_t) {
	listenerData := (*touchListenerData)(data)
	// Handle touch up event
	log.Printf("listenerData: %#v", listenerData)
}
