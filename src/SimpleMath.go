package main


type SimpleMath struct {
	x float64
	y float64
}

func (sm *SimpleMath) add() float64 {
	return sm.x + sm.y
}

func (sm *SimpleMath) sub() float64 {
	return sm.x - sm.y
}

func (sm *SimpleMath) mul() float64 {
	return sm.x * sm.y
}

func (sm *SimpleMath) div() float64 {
	return sm.x / sm.y
}
