package simulation

import (
	"context"
	"testing"
	"time"

	"evolab/internal/physics"
)

type stubRenderer struct {
	beforeCalls int
	afterCalls  int
}

func (s *stubRenderer) BeforeStep(step int, blobs []physics.Blob) {
	s.beforeCalls++
}

func (s *stubRenderer) AfterStep(step int, blobs []physics.Blob) {
	s.afterCalls++
}

func TestEngineRunsSteps(t *testing.T) {
	cfg := Config{Steps: 3, Timestep: time.Millisecond}
	blob := physics.NewBasicBlob("blob", physics.Vector{})
	renderer := &stubRenderer{}

	engine := NewEngine(cfg, []physics.Blob{blob}, renderer)
	if err := engine.Run(context.Background()); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if renderer.beforeCalls != cfg.Steps {
		t.Fatalf("expected %d before calls, got %d", cfg.Steps, renderer.beforeCalls)
	}

	if renderer.afterCalls != cfg.Steps {
		t.Fatalf("expected %d after calls, got %d", cfg.Steps, renderer.afterCalls)
	}

	if blob.Position() == (physics.Vector{}) {
		t.Fatalf("blob position did not update")
	}
}

func TestLoadConfigRequiresPath(t *testing.T) {
	if _, err := LoadConfig(""); err == nil {
		t.Fatal("expected error for empty config path")
	}
}
