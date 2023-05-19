package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lcallbacks
// #include "callbacks.h"
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

func initSeat(tinywl *tinywl) {
	tinywl.seat = C.wl_registry_bind(tinywl.registry, C.uint32_t(tinywl.seat.id), &C.wl_seat_interface, 1)
	C.wl_seat_add_listener(tinywl.seat, &C.seat_listener, unsafe.Pointer(tinywl))

	// Retrieve capabilities
	C.wl_pointer_add_listener(tinywl.pointer, &C.pointer_listener, unsafe.Pointer(tinywl))
	C.wl_keyboard_add_listener(tinywl.keyboard, &C.keyboard_listener, unsafe.Pointer(tinywl))
	C.wl_touch_add_listener(tinywl.touch, &C.touch_listener, unsafe.Pointer(tinywl))

	// Set seat capabilities
	C.wl_seat_set_capabilities(tinywl.seat, C.uint32_t(C.WL_SEAT_CAPABILITY_POINTER|C.WL_SEAT_CAPABILITY_KEYBOARD|C.WL_SEAT_CAPABILITY_TOUCH))
}
