package main

type INeuron interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{Neurons: make([]Neuron, count)}
}

func (n *NeuronLayer) Iter() []*Neuron {
	neurons := make([]*Neuron, 0, len(n.Neurons))
	for i := range n.Neurons {
		neurons = append(neurons, &n.Neurons[i])
	}
	return neurons
}

func Connect(left, right INeuron) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(5)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer2)
	Connect(layer1, neuron2)
	Connect(layer1, layer2)
}
