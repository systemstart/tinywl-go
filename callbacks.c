#include <wayland-client.h>
#include "_cgo_export.h"

// Seat Listener
void seatHandleCapabilities(void *data, struct wl_seat *seat, uint32_t capabilities) {
    seatHandleCapabilitiesCallback(data, seat, capabilities);
}

void seatHandleName(void *data, struct wl_seat *seat, const char *name) {
    seatHandleNameCallback(data, seat, name);
}

// Pointer Listener
void pointerHandleEnter(void *data, struct wl_pointer *pointer, uint32_t serial, struct wl_surface *surface, wl_fixed_t surface_x, wl_fixed_t surface_y) {
    pointerHandleEnterCallback(data, pointer, serial, surface, surface_x, surface_y);
}

void pointerHandleLeave(void *data, struct wl_pointer *pointer, uint32_t serial, struct wl_surface *surface) {
    pointerHandleLeaveCallback(data, pointer, serial, surface);
}

// Keyboard Listener
void keyboardHandleKeymap(void *data, struct wl_keyboard *keyboard, uint32_t format, int32_t fd, uint32_t size) {
    keyboardHandleKeymapCallback(data, keyboard, format, fd, size);
}

void keyboardHandleEnter(void *data, struct wl_keyboard *keyboard, uint32_t serial, struct wl_surface *surface, struct wl_array *keys) {
    keyboardHandleEnterCallback(data, keyboard, serial, surface, keys);
}

// Touch Listener
void touchHandleDown(void *data, struct wl_touch *touch, uint32_t serial, uint32_t time, struct wl_surface *surface, int32_t id, wl_fixed_t x, wl_fixed_t y) {
    touchHandleDownCallback(data, touch, serial, time, surface, id, x, y);
}

void touchHandleUp(void *data, struct wl_touch *touch, uint32_t serial, uint32_t time, int32_t id) {
    touchHandleUpCallback(data, touch, serial, time, id);
}
