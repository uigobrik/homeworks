package animal

type Cat struct {
	value AnimalStatus
}

func (c *Cat) GetName() string {
	return c.value.Name
}

func (c *Cat) GetAge() int {
	return c.value.Age
}

func (c *Cat) GetType() string {
	return c.value.Type
}

func (c *Cat) MakeSound() string {
	return "ニャオ"
}