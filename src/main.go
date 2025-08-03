package main

import (
	"fmt"

	extras "github.com/jb55matthews/gocpu/src/extras"
)

func main() {
	// clk := &Clock{}
	// sr := sequentiallogic.SRLatch{Q: false, NQ: true}
	// fmt.Printf("Q(f): %t, NQ(t): %t\n", sr.Q, sr.NQ)
	// sr.SetInputs(true, false)
	// fmt.Printf("Q(t): %t, NQ(f): %t\n", sr.Q, sr.NQ)
	// sr.SetInputs(false, false)
	// fmt.Printf("Q(prev): %t, NQ(prev): %t\n", sr.Q, sr.NQ)
	// sr.SetInputs(false, true)
	// fmt.Printf("Q(f): %t, NQ(t): %t\n", sr.Q, sr.NQ)
	// sr.SetInputs(false, false)
	// fmt.Printf("Q(prev): %t, NQ(prev): %t\n", sr.Q, sr.NQ)
	// sr.SetInputs(true, true)
	// fmt.Printf("Q(f): %t, NQ(f): %t\n", sr.Q, sr.NQ)
	// sr.SetInputs(true, false)
	// fmt.Printf("Q(t): %t, NQ(f): %t\n", sr.Q, sr.NQ)
	// sr.SetInputs(false, false)
	// fmt.Printf("Q(prev): %t, NQ(prev): %t\n", sr.Q, sr.NQ)

	// var dl sequentiallogic.DLatch
	// dl.Init()
	// fmt.Printf("Q(f): %t, NQ(t): %t\n", dl.Q, dl.NQ)
	// dl.SetInputs(true, true)
	// fmt.Printf("Q(t): %t, NQ(f): %t\n", dl.Q, dl.NQ)
	// dl.SetInputs(true, false)
	// fmt.Printf("Q(f): %t, NQ(t): %t\n", dl.Q, dl.NQ)
	// dl.SetInputs(true, true)
	// fmt.Printf("Q(t): %t, NQ(f): %t\n", dl.Q, dl.NQ)
	// dl.SetInputs(false, false)
	// fmt.Printf("Q(prev): %t, NQ(prev): %t\n", dl.Q, dl.NQ)
	// dl.SetInputs(false, true)
	// fmt.Printf("Q(prev): %t, NQ(prev): %t\n", dl.Q, dl.NQ)
	// dl.SetInputs(false, false)
	// fmt.Printf("Q(prev): %t, NQ(prev): %t\n", dl.Q, dl.NQ)


	// var df sequentiallogic.DFlipFlop
	// df.Init()
	// fmt.Println("Q(f) ", df.Q, " NQ(t), ", df.NQ)
	// df.SetInputs(false, true)
	// df.SetInputs(true, false)
	// fmt.Println("Q(prev) ", df.Q, " NQ(prev), ", df.NQ)
	// // df.SetInputs(false, false)
	// // fmt.Println("G(prev) ", df.Q, " NQ(prev), ", df.NQ)
	// df.SetInputs(false, true)
	// df.SetInputs(true, true)
	// fmt.Println("Q(prev) ", df.Q, " NQ(prev), ", df.NQ)


	// var r sequentiallogic.Register
	// r.Init()
	// fmt.Println("all false: ", r.Qs) 
	
	// var testbus, zerobus, onebus [32]bool
	// for i := range testbus {
	// 	if i%2 == 0 {
	// 		testbus[i] = true
	// 	}
	// 	onebus[i] = true
	// 	zerobus[i] = false
	// }

	// r.SetInputs(false, zerobus)
	// r.SetInputs(true, testbus)
	// fmt.Println("all false: ", r.Qs)
	// r.SetInputs(false, testbus)
	// r.SetInputs(true, r.Qs)
	// fmt.Println("testbus: ", r.Qs)


	d5 := extras.Decoder5{}
	d5.SetInputs([5]bool{false, false, false, false, false})
	fmt.Println("output: ", d5.OutputIndex())

}