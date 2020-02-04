package car

type Motorcycle struct {
	value CarValue
}

func (m *Motorcycle) Name() string {
	return m.value.Name
}

func (m *Motorcycle) MaxSpeed() int {
	return m.value.MaxSpeed
}

func (m *Motorcycle) TireCount() int {
	return m.value.TireCount
}

func NewMotorcycle() *Motorcycle {
	m := new(Motorcycle)
	m.value.Name = "Motorcycle"
	m.value.MaxSpeed = 210
	m.value.TireCount = 2
	return m
}
