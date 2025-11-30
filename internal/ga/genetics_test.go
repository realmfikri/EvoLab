package ga

import "testing"

func TestUniformCrossoverProducesCombinedGenome(t *testing.T) {
	parentA := NewMapGenome(map[string]float64{"speed": 1, "vision": 2})
	parentB := NewMapGenome(map[string]float64{"speed": 3, "agility": 4})

	crossover := UniformCrossover{}
	child := crossover.Crossover(parentA, parentB)

	if len(child.Traits()) == 0 {
		t.Fatal("expected child genome to contain traits")
	}

	if _, ok := child.Traits()["speed"]; !ok {
		t.Fatal("expected speed trait to be present")
	}
}

func TestGaussianMutationAltersTrait(t *testing.T) {
	genome := NewMapGenome(map[string]float64{"speed": 1})
	mutation := GaussianMutation{Scale: 0.5}

	mutated := mutation.Mutate(genome)
	if mutated.Traits()["speed"] == genome.Traits()["speed"] {
		t.Fatal("expected mutated trait to differ from original")
	}
}

func TestDenseNetworkForward(t *testing.T) {
	network := DenseNetwork{Weights: [][]float64{{1, 2}, {0.5, -0.5}}}
	outputs := network.Forward([]float64{2, 3})

	if len(outputs) != 2 {
		t.Fatalf("expected 2 outputs, got %d", len(outputs))
	}

	if outputs[0] != 8 || outputs[1] != -0.5 {
		t.Fatalf("unexpected output values: %v", outputs)
	}
}
