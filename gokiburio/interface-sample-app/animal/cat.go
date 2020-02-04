package animal

type CatValue struct {
	Value AnimalStatus
}

func (c *CatValue) GetName() string {
	return c.Value.Name
}

func (c *CatValue) GetAge() int {
	return c.Value.Age
}

func (c *CatValue) GetType() string {
	return c.Value.Type
}

func (c *CatValue) MakeSound() string {
	return "ニャオ"
}
