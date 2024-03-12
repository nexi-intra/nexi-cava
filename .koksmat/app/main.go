package main

import (
	"runtime/debug"
	"strings"

	"github.com/365admin/nexi-cava/magicapp"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: CAVA
description: Collection of stuff build for nexi "CAVA"
---

# CAVA - Meeting Management System`
	magicapp.Setup(".env")
	magicapp.RegisterServeCmd("CAVA", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.RegisterServiceCmd()
	magicapp.Execute(name, "CAVA", "")
}
