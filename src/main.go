package main
import (
	"fmt"
	latches "github.com/jb55matthews/gocpu/src/latches"
)

func main() {
	// clk := &Clock{}
	// sr := latches.SRLatch{Q: false, NQ: true}
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

	dl := latches.DLatch{Q: false, NQ: true}
	fmt.Printf("Q(f): %t, NQ(t): %t\n", dl.Q, dl.NQ)
	dl.SetInputs(true, true)
	fmt.Printf("Q(t): %t, NQ(f): %t\n", dl.Q, dl.NQ)
	dl.SetInputs(true, false)
	fmt.Printf("Q(f): %t, NQ(t): %t\n", dl.Q, dl.NQ)
	dl.SetInputs(true, true)
	fmt.Printf("Q(t): %t, NQ(f): %t\n", dl.Q, dl.NQ)
	dl.SetInputs(false, false)
	fmt.Printf("Q(prev): %t, NQ(prev): %t\n", dl.Q, dl.NQ)
	dl.SetInputs(false, true)
	fmt.Printf("Q(prev): %t, NQ(prev): %t\n", dl.Q, dl.NQ)
	dl.SetInputs(false, false)
	fmt.Printf("Q(prev): %t, NQ(prev): %t\n", dl.Q, dl.NQ)

}