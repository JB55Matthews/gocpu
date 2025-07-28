package sequentiallogic

import (gates "github.com/jb55matthews/gocpu/src/gates"
"fmt")

type SRLatch struct {
    nor1, nor2 gates.NorGate
    Q, NQ        bool
}

func (l *SRLatch) SetInputs(S, R bool) {
    // fmt.Println("s, r" ,S, R)
    // fmt.Println("q, nq", l.Q, l.NQ)
    l.nor1.SetInputs(R, l.NQ)
    l.Q = l.nor1.Output()

    l.nor2.SetInputs(S, l.Q)
    l.NQ = l.nor2.Output()

	l.nor1.SetInputs(R, l.NQ)
    l.Q = l.nor1.Output()
    fmt.Printf("")
    // fmt.Println("q, nq", l.Q, l.NQ)
	
}

// main init to make sure we don't have q=f, nq=f in srlatch
func (l *SRLatch) Init(){
	l.Q = false
	l.NQ = true
}