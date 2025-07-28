package main

type ClockedComponent interface {
	Tick(clk bool)
}

type Clock struct {
	clk bool
	components []ClockedComponent

}

func (c *Clock) AddComponent(component ClockedComponent) {
    c.components = append(c.components, component)
}

func (c *Clock) Tick() {
	c.clk = !c.clk
	for _, component := range c.components {
		component.Tick(c.clk)
	}
}