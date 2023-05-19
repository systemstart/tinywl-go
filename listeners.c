#include <wayland-client.h>
#include "_cgo_export.h"

// Registry Listener
void registryHandleGlobal(void *data, struct wl_registry *registry, uint32_t name, const char *interface, uint32_t version) {
    registryHandleGlobalCallback(data, registry, name, interface, version);
}

void registryHandleGlobalRemove(void *data, struct wl_registry *registry, uint32_t name) {
    registryHandleGlobalRemoveCallback(data, registry, name);
}

// Shell Surface Listener
void shellSurfaceHandlePing(void *data, struct wl_shell_surface *shell_surface, uint32_t serial) {
    shellSurfaceHandlePingCallback(data, shell_surface, serial);
}

void shellSurfaceHandleConfigure(void *data, struct wl_shell_surface *shell_surface, uint32_t edges, int32_t width, int32_t height) {
    shellSurfaceHandleConfigureCallback(data, shell_surface, edges, width, height);
}

void shellSurfaceHandlePopupDone(void *data, struct wl_shell_surface *shell_surface) {
    shellSurfaceHandlePopupDoneCallback(data, shell_surface);
}
