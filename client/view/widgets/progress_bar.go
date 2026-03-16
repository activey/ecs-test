package widgets

import "github.com/yohamta/furex/v2"

type ProgressBar struct {
	Progress int

	view        *furex.View
	fillHandler *ProgressBarFillHandler
}

func NewProgressBar(progress int) *ProgressBar {
	progressBarView := &furex.View{
		Direction:    furex.Column,
		AlignContent: furex.AlignContentStart,
		AlignItems:   furex.AlignItemStart,
		Width:        32 * 11,
		Height:       32,
		Handler:      NewProgressBarHandler(),
	}

	fillHandler := NewProgressBarFillHandler(progress)
	progressBarView.AddChild(&furex.View{
		Handler: fillHandler,
		Height:  16,
		Width:   32 * 11,
	})

	fillHandler.UpdateProgress(progress)
	return &ProgressBar{
		Progress:    progress,
		view:        progressBarView,
		fillHandler: fillHandler,
	}
}

func (p *ProgressBar) UpdateProgress(newProgress int) {
	p.fillHandler.UpdateProgress(newProgress)
}

func (p *ProgressBar) View() *furex.View {
	return p.view
}
