package ggcore

type Timestep float32

func (timestep Timestep) GetSeconds() float32 {
	return float32(timestep)
}

func (timestep Timestep) GetMilliSeconds() float32 {
	return float32(timestep) * 1000
}
