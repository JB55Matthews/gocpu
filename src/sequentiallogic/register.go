package sequentiallogic

type Register struct {
	flipflops [32]DFlipFlop
	Qs [32]bool
}

func (r *Register) SetInputs(clk bool, bus [32]bool) {
	for i := range bus {
		r.flipflops[i].SetInputs(clk, bus[i])
		r.Qs[i] = r.flipflops[i].Q
	}
}

func (r* Register) Init() {
	for i, f := range r.flipflops {
		f.Init()
		r.Qs[i] = f.Q
	}
}