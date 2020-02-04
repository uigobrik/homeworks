package car

type NormalCar struct {
	value CarValue
}

func (n *NormalCar) Name() string {
	return n.value.Name
}

func (n *NormalCar) MaxSpeed() int {
	return n.value.MaxSpeed
}

func (n *NormalCar) TireCount() int {
	return n.value.TireCount
}

func NewNormalCar() *NormalCar {
	n := new(NormalCar)
	n.value.Name = "NormaCar"
	n.value.MaxSpeed = 320
	n.value.TireCount = 4
	return n
}
