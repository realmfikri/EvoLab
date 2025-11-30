package main

import (
	"context"
	"log"
	"time"

	"evolab/internal/physics"
	"evolab/internal/simulation"
)

func main() {
	ctx := context.Background()

	cfg, err := simulation.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	blobs := []physics.Blob{
		physics.NewBasicBlob("seed-1", physics.Vector{X: 0, Y: 0}),
	}

	engine := simulation.NewEngine(cfg, blobs, simulation.NoOpRenderer{})
	if err := engine.Run(ctx); err != nil {
		log.Fatalf("simulation terminated with error: %v", err)
	}

	time.Sleep(10 * time.Millisecond)
}
