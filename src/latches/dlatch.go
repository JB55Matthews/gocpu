package latches

import gates "github.com/jb55matthews/gocpu/src/gates"

type DLatch struct {
    and1, and2 gates.AndGate
	notD gates.NotGate
	srlatch SRLatch
    Q, NQ bool
}

func (l *DLatch) SetInputs(clk, d bool) {
	l.notD.SetInputs(d)
	nd := l.notD.Output()

	l.and1.SetInputs(clk, nd)
	r := l.and1.Output()

	l.and2.SetInputs(clk, d)
	s := l.and2.Output()

	l.srlatch.SetInputs(s, r)
	
	l.Q, l.NQ = l.srlatch.Q, l.srlatch.NQ
}