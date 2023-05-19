package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lcallbacks -llisteners
// #cgo !windows CFLAGS: -DXKB_COMPOSE -DXKB_TEXT -DXKB_ATOMS -DXKB_VMOD -DXKB_MODIFIERS -DXKB_GLIB -DXKB_KEYMAP_FLAGS
// #cgo !windows LDFLAGS: -lxkbcommon
// #include "callbacks.h"
// #include "listeners.h"
// #include <xkbcommon/xkbcommon.h>
//
// // Manually define xkb_mod_index_t to match the typedef in the C header file
// typedef uint32_t xkb_mod_index_t;
import "C"

/*

#cgo pkg-config: wayland-client wayland-egl wayland-cursor xkbcommon egl glesv2
#include <wayland-client.h>
#include <wayland-egl.h>
#include <wayland-cursor.h>
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

// Forward declarations for C functions
static void handle_ping(void *data, struct wl_shell_surface *shell_surface, uint32_t serial);
*/
import "C"

import (
	"log"
	"unsafe"
)

// Pointer Listener
type pointerListenerData struct {
	tinywl *tinywl
}

//export handlePing
func handlePing(data unsafe.Pointer, shellSurface *C.struct_wl_shell_surface, serial C.uint32_t) {
	C.wl_shell_surface_pong(shellSurface, serial)
}

func main() {
	t := newTinyWL()

	display := C.wl_display_connect(nil)
	if display == nil {
		log.Fatal("Failed to connect to Wayland display")
	}
	t.display = display

	t.registry = C.wl_display_get_registry(display)
	initRegistry(t)

	C.wl_display_roundtrip(display)

	// Make sure the compositor supports the required interfaces
	if !t.compositor || !t.shell {
		log.Fatal("Compositor or shell interface not available")
	}

	t.surface = C.wl_compositor_create_surface(t.compositor)
	initShellSurface(t, t.surface)

	C.wl_display_roundtrip(display)

	// Main event loop
	for C.wl_display_dispatch(display) != -1 {
		// Continue dispatching events
	}

	// Cleanup resources
	C.wl_shell_surface_destroy(t.shellSurface)
	C.wl_surface_destroy(t.surface)
	C.wl_registry_destroy(t.registry)
	C.wl_display_disconnect(t.display)
}
