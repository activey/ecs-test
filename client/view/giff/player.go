package giff

import (
	"time"
)

type GIFPlayer struct {
	delays       []int // Delay for each frame in milliseconds
	currentDelay int   // Current frame delay
	speed        float64
	elapsed      time.Duration // Time elapsed since last frame update
	lastUpdate   time.Time     // Last time frame was updated
	frameIndex   int           // Current frame index
	observers    []func()
}

func NewPlayer(i *GIFImage, speed float64) *GIFPlayer {
	return &GIFPlayer{
		delays:       i.DelayMilliSec(),    // Use delays for all frames
		currentDelay: i.DelayMilliSec()[0], // Initial delay
		speed:        speed,
		observers:    make([]func(), 0),
		lastUpdate:   time.Now(),
	}
}

func (p *GIFPlayer) AddObserver(f func()) {
	p.observers = append(p.observers, f)
}

func (p *GIFPlayer) Update() {
	now := time.Now()
	p.elapsed += now.Sub(p.lastUpdate) // Accumulate elapsed time
	p.lastUpdate = now

	waitDuration := p.calcWaitDuration()

	// If enough time has passed, trigger frame change
	if p.elapsed >= waitDuration {
		p.elapsed -= waitDuration // Reduce elapsed time by the frame duration

		// Update the frame index and notify observers
		p.frameIndex = (p.frameIndex + 1) % len(p.delays) // Wrap around the frame index
		p.currentDelay = p.delays[p.frameIndex]           // Update the delay for the current frame
		for _, o := range p.observers {
			o() // Notify observers of frame change
		}
	}
}

func (p *GIFPlayer) calcWaitDuration() time.Duration {
	ms := float64(p.currentDelay) * float64(time.Millisecond) / p.speed
	return time.Duration(ms)
}
