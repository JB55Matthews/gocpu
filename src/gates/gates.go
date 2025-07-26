package gates

type Gate interface {
	Output() bool
}

type AndGate struct{
	inA, inB bool
	output bool
}

func (g *AndGate) SetInputs(a, b bool){
	g.inA = a
	g.inB = b
	g.output = g.inA && g.inB
}

func (g *AndGate) Output() bool {
	return g.output
}

type NandGate struct{
	inA, inB bool
	output bool
}

func (g *NandGate) SetInputs(a, b bool){
	g.inA = a
	g.inB = b
	g.output = !(g.inA && g.inB)
}

func (g *NandGate) Output() bool {
	return g.output
}

type OrGate struct{
	inA, inB bool
	output bool
}

func (g *OrGate) SetInputs(a, b bool){
	g.inA = a
	g.inB = b
	g.output = g.inA || g.inB
}

func (g *OrGate) Output() bool {
	return g.output
}

type NorGate struct{
	inA, inB bool
	output bool
}

func (g *NorGate) SetInputs(a, b bool){
	g.inA = a
	g.inB = b
	g.output = !(g.inA || g.inB)
}

func (g *NorGate) Output() bool {
	return g.output
}

type NotGate struct{
	inA bool
	output bool
}

func (g *NotGate) SetInputs(a bool){
	g.inA = a
	g.output = !g.inA
}

func (g *NotGate) Output() bool {
	return g.output
}