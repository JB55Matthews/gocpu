package sequentiallogic

import (gates "github.com/jb55matthews/gocpu/src/gates" 
"fmt")

type DFlipFlop struct {
	masterLatch, slaveLatch DLatch
	notClk gates.NotGate
	Q, NQ bool
}

func (f *DFlipFlop) SetInputs(clk, d bool) {
	f.notClk.SetInputs(clk)
	nclk := f.notClk.Output()

	f.masterLatch.SetInputs(nclk, d)
	// fmt.Println("nclk, d, masterQ", nclk, d, f.masterLatch.Q)

	// fmt.Println("slaveQ", f.slaveLatch.Q)
	f.slaveLatch.SetInputs(clk, f.masterLatch.Q)
	// fmt.Println("clk, masterQ, slaveQ", clk, f.masterLatch.Q, f.slaveLatch.Q)
	fmt.Printf("")
	f.Q, f.NQ = f.slaveLatch.Q, f.slaveLatch.NQ
}

func (f *DFlipFlop) Init(){
	f.masterLatch.Init()
	f.slaveLatch.Init()
	f.Q, f.NQ = f.slaveLatch.Q, f.slaveLatch.NQ
}