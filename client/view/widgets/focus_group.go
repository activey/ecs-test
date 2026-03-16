package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/furex/v2"
)

type Focusable interface {
	Focus()
	Blur()
}

type FocusGroup struct {
	items        []Focusable
	focusedIndex int
}

func NewFocusGroup() *FocusGroup {
	return &FocusGroup{
		items: make([]Focusable, 0),
	}
}

func (f *FocusGroup) AddItem(focusable Focusable) {
	f.items = append(f.items, focusable)

	if len(f.items) == 1 {
		f.items[0].Focus()
	}
}

func (f *FocusGroup) Update(v *furex.View) {
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		f.focusedIndex++
		if f.focusedIndex > len(f.items)-1 {
			f.focusedIndex = 0
		}

		for i, item := range f.items {
			if i == f.focusedIndex {
				item.Focus()
			} else {
				item.Blur()
			}
		}
	}
}

func (f *FocusGroup) Focused(focusable Focusable) {
	for i, item := range f.items {
		if item == focusable {
			f.focusedIndex = i
		} else {
			item.Blur()
		}
	}
}

func (f *FocusGroup) Reset() {
	f.focusedIndex = 0

	for i, item := range f.items {
		if i == f.focusedIndex {
			item.Focus()
		} else {
			item.Blur()
		}
	}
}
