package component

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/furex/v2"
)

type DebugData struct {
	debugMode bool
}

func (d *DebugData) IsEnabled() bool {
	return d.debugMode
}

func (d *DebugData) EnableDebug() {
	d.debugMode = true

	// should it be here? ;)
	furex.Debug = true
}

func (d *DebugData) DisableDebug() {
	d.debugMode = false
	furex.Debug = false
}

func (d *DebugData) ToggleDebug() {
	d.debugMode = !d.debugMode
	furex.Debug = d.debugMode
}

var Debug = donburi.NewComponentType[DebugData]()
