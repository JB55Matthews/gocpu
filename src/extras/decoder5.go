package extras

import gates "github.com/jb55matthews/gocpu/src/gates"

type Decoder5 struct {
	nadr [5]gates.NotGate
	ands [32]AndGate5
	Output [32]bool
}

func (d5 *Decoder5) SetInputs(adr [5]bool) {
	for i, bit := range adr {
		d5.nadr[i].SetInputs(bit)
	}

	d5.ands[0].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), d5.nadr[2].Output(), d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[0] = d5.ands[0].Output()

	d5.ands[1].SetInputs(adr[0], d5.nadr[1].Output(), d5.nadr[2].Output(), d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[1] = d5.ands[1].Output()

	d5.ands[2].SetInputs(d5.nadr[0].Output(), adr[1], d5.nadr[2].Output(), d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[2] = d5.ands[2].Output()

	d5.ands[3].SetInputs(adr[0], adr[1], d5.nadr[2].Output(), d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[3] = d5.ands[3].Output()

	d5.ands[4].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), adr[2], d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[4] = d5.ands[4].Output()

	d5.ands[5].SetInputs(adr[0], d5.nadr[1].Output(), adr[2], d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[5] = d5.ands[5].Output()

	d5.ands[6].SetInputs(d5.nadr[0].Output(), adr[1], adr[2], d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[6] = d5.ands[6].Output()

	d5.ands[7].SetInputs(adr[0], adr[1], adr[2], d5.nadr[3].Output(), d5.nadr[4].Output())
	d5.Output[7] = d5.ands[7].Output()

	d5.ands[8].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), d5.nadr[2].Output(), adr[3], d5.nadr[4].Output())
	d5.Output[8] = d5.ands[8].Output()

	d5.ands[9].SetInputs(adr[0], d5.nadr[1].Output(), d5.nadr[2].Output(), adr[3], d5.nadr[4].Output())
	d5.Output[9] = d5.ands[9].Output()

	d5.ands[10].SetInputs(d5.nadr[0].Output(), adr[1], d5.nadr[2].Output(), adr[3], d5.nadr[4].Output())
	d5.Output[10] = d5.ands[10].Output()

	d5.ands[11].SetInputs(adr[0], adr[1], d5.nadr[2].Output(), adr[3], d5.nadr[4].Output())
	d5.Output[11] = d5.ands[11].Output()

	d5.ands[12].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), adr[2], adr[3], d5.nadr[4].Output())
	d5.Output[12] = d5.ands[12].Output()

	d5.ands[13].SetInputs(adr[0], d5.nadr[1].Output(), adr[2], adr[3], d5.nadr[4].Output())
	d5.Output[13] = d5.ands[13].Output()

	d5.ands[14].SetInputs(d5.nadr[0].Output(), adr[1], adr[2], adr[3], d5.nadr[4].Output())
	d5.Output[14] = d5.ands[14].Output()

	d5.ands[15].SetInputs(adr[0], adr[1], adr[2], adr[3], d5.nadr[4].Output())
	d5.Output[15] = d5.ands[15].Output()

	d5.ands[16].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), d5.nadr[2].Output(), d5.nadr[3].Output(), adr[4])
	d5.Output[16] = d5.ands[16].Output()

	d5.ands[17].SetInputs(adr[0], d5.nadr[1].Output(), d5.nadr[2].Output(), d5.nadr[3].Output(), adr[4])
	d5.Output[17] = d5.ands[17].Output()

	d5.ands[18].SetInputs(d5.nadr[0].Output(), adr[1], d5.nadr[2].Output(), d5.nadr[3].Output(), adr[4])
	d5.Output[18] = d5.ands[18].Output()

	d5.ands[19].SetInputs(adr[0], adr[1], d5.nadr[2].Output(), d5.nadr[3].Output(), adr[4])
	d5.Output[19] = d5.ands[19].Output()

	d5.ands[20].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), adr[2], d5.nadr[3].Output(), adr[4])
	d5.Output[20] = d5.ands[20].Output()

	d5.ands[21].SetInputs(adr[0], d5.nadr[1].Output(), adr[2], d5.nadr[3].Output(), adr[4])
	d5.Output[21] = d5.ands[21].Output()

	d5.ands[22].SetInputs(d5.nadr[0].Output(), adr[1], adr[2], d5.nadr[3].Output(), adr[4])
	d5.Output[22] = d5.ands[22].Output()

	d5.ands[23].SetInputs(adr[0], adr[1], adr[2], d5.nadr[3].Output(), adr[4])
	d5.Output[23] = d5.ands[23].Output()

	d5.ands[24].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), d5.nadr[2].Output(), adr[3], adr[4])
	d5.Output[24] = d5.ands[24].Output()

	d5.ands[25].SetInputs(adr[0], d5.nadr[1].Output(), d5.nadr[2].Output(), adr[3], adr[4])
	d5.Output[25] = d5.ands[25].Output()

	d5.ands[26].SetInputs(d5.nadr[0].Output(), adr[1], d5.nadr[2].Output(), adr[3], adr[4])
	d5.Output[26] = d5.ands[26].Output()

	d5.ands[27].SetInputs(adr[0], adr[1], d5.nadr[2].Output(), adr[3], adr[4])
	d5.Output[27] = d5.ands[27].Output()

	d5.ands[28].SetInputs(d5.nadr[0].Output(), d5.nadr[1].Output(), adr[2], adr[3], adr[4])
	d5.Output[28] = d5.ands[28].Output()

	d5.ands[29].SetInputs(adr[0], d5.nadr[1].Output(), adr[2], adr[3], adr[4])
	d5.Output[29] = d5.ands[29].Output()

	d5.ands[30].SetInputs(d5.nadr[0].Output(), adr[1], adr[2], adr[3], adr[4])
	d5.Output[30] = d5.ands[30].Output()

	d5.ands[31].SetInputs(adr[0], adr[1], adr[2], adr[3], adr[4])
	d5.Output[31] = d5.ands[31].Output()

}

// Used to get index in RegisterFile a convience. 
// We could make same maual 32 bit selection in RF a we do above,
// however this is just unnecessary, so we do this
func (d5 *Decoder5) OutputIndex() int {
	var result int = 0
	for i, b := range d5.Output {
		if b {
			result = i
			break
		}
	}
	return result
}


type AndGate5 struct{
	a0, a1, a2, a3, a4 bool
	output bool
}

func (g *AndGate5) SetInputs(a, b, c, d, e bool){
	g.a0, g.a1, g.a2, g.a3, g.a4 = a, b,c, d, e
	g.output = g.a0 && g.a1 && g.a2 && g.a3 && g.a4
}

func (g *AndGate5) Output() bool {
	return g.output
}