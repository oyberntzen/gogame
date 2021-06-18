package ggcore

type LayerStack struct {
	Layers           []Layer
	layerInsertIndex int
}

func (stack *LayerStack) Delete() {
	for _, layer := range stack.Layers {
		layer.OnDetach()
	}
}

func (stack *LayerStack) PushLayer(layer Layer) {
	stack.Layers = append(stack.Layers, nil)
	copy(stack.Layers[stack.layerInsertIndex+1:], stack.Layers[stack.layerInsertIndex:])
	stack.Layers[stack.layerInsertIndex] = layer
	stack.layerInsertIndex++

	layer.OnAttach()
}

func (stack *LayerStack) PushOverlay(layer Layer) {
	stack.Layers = append(stack.Layers, layer)

	layer.OnAttach()
}

func (stack *LayerStack) PopLayer(layer Layer) {
	for i := 0; i < stack.layerInsertIndex; i++ {
		if layer == stack.Layers[i] {
			stack.Layers = append(stack.Layers[:i], stack.Layers[i+1:]...)
			layer.OnDetach()
			stack.layerInsertIndex--
			break
		}
	}
}

func (stack *LayerStack) PopOverlay(layer Layer) {
	for i := stack.layerInsertIndex; i < len(stack.Layers); i++ {
		if layer == stack.Layers[i] {
			stack.Layers = append(stack.Layers[:i], stack.Layers[i+1:]...)
			layer.OnDetach()
			break
		}
	}
}
