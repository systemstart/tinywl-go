package main

/*
#cgo pkg-config: wayland-client wayland-egl wayland-cursor xkbcommon egl glesv2
#include <wayland-client.h>
#include <wayland-egl.h>
#include <wayland-cursor.h>
#include <xkbcommon/xkbcommon.h>
#include <EGL/egl.h>
#include <EGL/eglext.h>
#include <GLES2/gl2.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>

// Placeholder definitions to satisfy the Go compiler.
// Replace with the actual struct definitions when translating further.
typedef struct wl_seat wl_seat;
typedef struct wl_pointer wl_pointer;
typedef struct wl_keyboard wl_keyboard;
typedef struct wl_touch wl_touch;
typedef struct wl_surface wl_surface;
typedef struct wl_shell_surface wl_shell_surface;
typedef struct wl_egl_window wl_egl_window;
typedef struct wl_registry wl_registry;
typedef struct wl_compositor wl_compositor;
typedef struct wl_shell wl_shell;
typedef struct wl_shm wl_shm;
typedef struct wl_cursor_theme wl_cursor_theme;
typedef struct wl_cursor wl_cursor;
typedef struct xkb_context xkb_context;
typedef struct xkb_keymap xkb_keymap;
typedef struct xkb_state xkb_state;
typedef int xkb_mod_index_t;

// Forward declarations for C functions
static void handle_ping(void *data, struct wl_shell_surface *shell_surface, uint32_t serial);
*/
import "C"

import (
	"log"
	"unsafe"
)

type seat struct {
	seat          *C.struct_wl_seat
	pointer       *C.struct_wl_pointer
	keyboard      *C.struct_wl_keyboard
	touch         *C.struct_wl_touch
	cursorSurface *C.struct_wl_surface
}

var seat_listener = C.struct_wl_seat_listener{
	capabilities: (*[0]byte)(C.seatHandleCapabilities),
	name:         (*[0]byte)(C.seatHandleName),
}

// Pointer Listener
type pointerListenerData struct {
	tinywl *tinywl
}

func pointerHandleEnter(data unsafe.Pointer, pointer *C.struct_wl_pointer, serial C.uint32_t, surface *C.struct_wl_surface, surface_x C.wl_fixed_t, surface_y C.wl_fixed_t) {
	listenerData := (*pointerListenerData)(data)
	// Handle pointer enter event
	log.Printf("pointerHandleEnter: %#v", listenerData)
}

func pointerHandleLeave(data unsafe.Pointer, pointer *C.struct_wl_pointer, serial C.uint32_t, surface *C.struct_wl_surface) {
	listenerData := (*pointerListenerData)(data)
	// Handle pointer leave event
	log.Printf("pointerHandleLeave: %#v", listenerData)
}

//export handlePing
func handlePing(data unsafe.Pointer, shellSurface *C.struct_wl_shell_surface, serial C.uint32_t) {
	C.wl_shell_surface_pong(shellSurface, serial)
}

func main() {
	tinywl := newTinyWL()

	display := C.wl_display_connect(nil)
	if display == nil {
		log.Fatal("Failed to connect to Wayland display")
	}
	tinywl.display = display

	tinywl.registry = C.wl_display_get_registry(display)
	initRegistry(tinywl)

	C.wl_display_roundtrip(display)

	// Make sure the compositor supports the required interfaces
	if !tinywl.compositor || !tinywl.shell {
		log.Fatal("Compositor or shell interface not available")
	}

	tinywl.surface = C.wl_compositor_create_surface(tinywl.compositor)
	initShellSurface(tinywl, tinywl.surface)

	C.wl_display_roundtrip(display)

	// Main event loop
	for C.wl_display_dispatch(display) != -1 {
		// Continue dispatching events
	}

	// Cleanup resources
	C.wl_shell_surface_destroy(tinywl.shellSurface)
	C.wl_surface_destroy(tinywl.surface)
	C.wl_registry_destroy(tinywl.registry)
	C.wl_display_disconnect(tinywl.display)
}
