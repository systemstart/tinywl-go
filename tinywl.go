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
	tinywl := &tinywl{}

	// Initialize fields to nil or zero values
	tinywl.display = nil
	tinywl.registry = nil
	tinywl.compositor = nil
	tinywl.shell = nil
	tinywl.shellSurface = nil
	tinywl.surface = nil

	return tinywl
}
