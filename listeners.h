#ifndef LISTENERS_H
#define LISTENERS_H

#include <wayland-client.h>

// Registry Listener
void registryHandleGlobal(void *data, struct wl_registry *registry, uint32_t name, const char *interface, uint32_t version);
void registryHandleGlobalRemove(void *data, struct wl_registry *registry, uint32_t name);

// Shell Surface Listener
void shellSurfaceHandlePing(void *data, struct wl_shell_surface *shell_surface, uint32_t serial);
void shellSurfaceHandleConfigure(void *data, struct wl_shell_surface *shell_surface, uint32_t edges, int32_t width, int32_t height);
void shellSurfaceHandlePopupDone(void *data, struct wl_shell_surface *shell_surface);

#endif
