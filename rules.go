package controlledczech

import "embed"

// Rules obsahuje strojově čitelné části specifikace používané linterem.
//
//go:embed pravidla/*
var Rules embed.FS
