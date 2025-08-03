package main

import (
	logicunits "github.com/jb55matthews/gocpu/src/sequentiallogic"
	extras "github.com/jb55matthews/gocpu/src/extras"
	gates "github.com/jb55matthews/gocpu/src/gates"
)

type RegisterFile struct {
	registers [32]logicunits.Register
	decoder extras.Decoder5
	wd3, rd1, rd2 [32] bool
	we3 bool

	andwe3 gates.AndGate
	
}

func (rf *RegisterFile) SetInputs(a1, a2, a3 [5]bool, we3, clk bool, wd3 [32]bool) {
	rf.decoder.SetInputs(a1)
	rf.rd1 = rf.registers[rf.decoder.OutputIndex()].Qs

	rf.decoder.SetInputs(a2)
	rf.rd2 = rf.registers[rf.decoder.OutputIndex()].Qs

	rf.wd3 = wd3
	rf.we3 = we3

	
}

func (rf *RegisterFile) Init() {
	for _, r := range rf.registers {
		r.Init()
	}	
}