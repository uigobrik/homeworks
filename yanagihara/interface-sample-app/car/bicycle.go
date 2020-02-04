package car

type Bicycle struct {
	value CarValue
}

func (b *Bicycle) Name() string {
	return b.value.Name
}

func (b *Bicycle) MaxSpeed() int {
	return b.value.MaxSpeed
}

func (b *Bicycle) TireCount() int {
	return b.value.TireCount
}

func NewBicycle() *Bicycle {
	b := new(Bicycle)
	b.value.Name = "Bicycle"
	b.value.MaxSpeed = 60
	b.value.TireCount = 2
	return b
}
