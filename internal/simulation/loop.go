package simulation

import (
	"context"
	"fmt"
	"time"

	"evolab/internal/physics"
)

// Renderer allows instrumentation of the simulation loop without coupling the
// core to a specific output technology.
type Renderer interface {
	BeforeStep(step int, blobs []physics.Blob)
	AfterStep(step int, blobs []physics.Blob)
}

// NoOpRenderer provides a default renderer that performs no work.
type NoOpRenderer struct{}

// BeforeStep is a no-op implementation.
func (NoOpRenderer) BeforeStep(step int, blobs []physics.Blob) {}

// AfterStep is a no-op implementation.
func (NoOpRenderer) AfterStep(step int, blobs []physics.Blob) {}

// Engine coordinates the simulation loop.
type Engine struct {
	cfg      Config
	blobs    []physics.Blob
	renderer Renderer
}

// NewEngine constructs an Engine with the provided configuration, blobs, and
// renderer.
func NewEngine(cfg Config, blobs []physics.Blob, renderer Renderer) *Engine {
	if renderer == nil {
		renderer = NoOpRenderer{}
	}

	return &Engine{cfg: cfg, blobs: blobs, renderer: renderer}
}

// Run executes the simulation for the configured number of steps.
func (e *Engine) Run(ctx context.Context) error {
	if e.cfg.Steps <= 0 {
		return fmt.Errorf("invalid step count: %d", e.cfg.Steps)
	}

	for step := 0; step < e.cfg.Steps; step++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e.renderer.BeforeStep(step, e.blobs)
		e.stepOnce(e.cfg.Timestep)
		e.renderer.AfterStep(step, e.blobs)
	}

	return nil
}

func (e *Engine) stepOnce(delta time.Duration) {
	for _, blob := range e.blobs {
		blob.Update(delta)
	}
}
