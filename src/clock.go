package main

type ClockedComponent interface {
	Tick(ishigh bool)
}

type Clock struct {
	ishigh bool
	components []ClockedComponent

}

func (c *Clock) AddComponent(component ClockedComponent) {
    c.components = append(c.components, component)
}

func (c *Clock) Tick() {
	c.ishigh = !c.ishigh
	for _, component := range c.components {
		component.Tick(c.ishigh)
	}
}