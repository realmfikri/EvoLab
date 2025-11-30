package ga

import "math/rand"

// Genome describes the genetic payload for an entity.
type Genome interface {
	Traits() map[string]float64
}

// CrossoverStrategy merges parent genomes into a new genome.
type CrossoverStrategy interface {
	Crossover(a, b Genome) Genome
}

// MutationStrategy mutates a genome in place or returns a new copy.
type MutationStrategy interface {
	Mutate(g Genome) Genome
}

// MapGenome is a minimal genome implementation backed by a map of traits.
type MapGenome struct {
	traits map[string]float64
}

// NewMapGenome constructs a MapGenome with the provided traits.
func NewMapGenome(traits map[string]float64) *MapGenome {
	return &MapGenome{traits: traits}
}

// Traits returns the trait map for the genome.
func (g *MapGenome) Traits() map[string]float64 { return g.traits }

// UniformCrossover mixes traits from two genomes at random.
type UniformCrossover struct{}

// Crossover returns a new genome with randomly selected traits from each parent.
func (UniformCrossover) Crossover(a, b Genome) Genome {
	child := make(map[string]float64)
	for key, value := range a.Traits() {
		child[key] = value
	}

	for key, value := range b.Traits() {
		if rand.Float64() > 0.5 {
			child[key] = value
		}
	}

	return &MapGenome{traits: child}
}

// GaussianMutation applies a small random delta to each trait.
type GaussianMutation struct {
	Scale float64
}

// Mutate returns a new genome with perturbed traits.
func (m GaussianMutation) Mutate(g Genome) Genome {
	mutated := make(map[string]float64)
	for key, value := range g.Traits() {
		mutated[key] = value + rand.NormFloat64()*m.Scale
	}
	return &MapGenome{traits: mutated}
}
