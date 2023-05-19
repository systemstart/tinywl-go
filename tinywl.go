package main

import "C"

type tinywl struct {
	display      *C.struct_wl_display
	registry     *C.struct_wl_registry
	compositor   *C.struct_wl_compositor
	shell        *C.struct_wl_shell
	shellSurface *C.struct_wl_shell_surface
	surface      *C.struct_wl_surface
}

func newTinyWL() *tinywl {
	t := &tinywl{}

	// Initialize fields to nil or zero values
	t.display = nil
	t.registry = nil
	t.compositor = nil
	t.shell = nil
	t.shellSurface = nil
	t.surface = nil

	return t
}
