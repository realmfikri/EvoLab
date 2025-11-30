package ga

// NeuralNetwork defines the behavior of a minimal feed-forward network.
type NeuralNetwork interface {
	Forward(inputs []float64) []float64
}

// DenseNetwork is a minimal implementation that performs a single dense layer
// multiply without activation.
type DenseNetwork struct {
	Weights [][]float64
}

// Forward multiplies inputs by the weight matrix producing an output vector.
func (n DenseNetwork) Forward(inputs []float64) []float64 {
	outputs := make([]float64, len(n.Weights))
	for i, row := range n.Weights {
		var sum float64
		for j, weight := range row {
			if j < len(inputs) {
				sum += weight * inputs[j]
			}
		}
		outputs[i] = sum
	}
	return outputs
}
