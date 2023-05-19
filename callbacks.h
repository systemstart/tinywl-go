#ifndef CALLBACKS_H
#define CALLBACKS_H

#include <wayland-client.h>

// Seat Listener
void seatHandleCapabilities(void *data, struct wl_seat *seat, uint32_t capabilities);
void seatHandleName(void *data, struct wl_seat *seat, const char *name);

// Pointer Listener
void pointerHandleEnter(void *data, struct wl_pointer *pointer, uint32_t serial, struct wl_surface *surface, wl_fixed_t surface_x, wl_fixed_t surface_y);
void pointerHandleLeave(void *data, struct wl_pointer *pointer, uint32_t serial, struct wl_surface *surface);

// Keyboard Listener
void keyboardHandleKeymap(void *data, struct wl_keyboard *keyboard, uint32_t format, int32_t fd, uint32_t size);
void keyboardHandleEnter(void *data, struct wl_keyboard *keyboard, uint32_t serial, struct wl_surface *surface, struct wl_array *keys);

// Touch Listener
void touchHandleDown(void *data, struct wl_touch *touch, uint32_t serial, uint32_t time, struct wl_surface *surface, int32_t id, wl_fixed_t x, wl_fixed_t y);
void touchHandleUp(void *data, struct wl_touch *touch, uint32_t serial, uint32_t time, int32_t id);

#endif
