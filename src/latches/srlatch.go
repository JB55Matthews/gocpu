package latches

import gates "github.com/jb55matthews/gocpu/src/gates"

type SRLatch struct {
    nor1, nor2 gates.NorGate
    Q, NQ        bool
}

func (l *SRLatch) SetInputs(S, R bool) {

    l.nor1.SetInputs(R, l.NQ)
    l.Q = l.nor1.Output()

    l.nor2.SetInputs(S, l.Q)
    l.NQ = l.nor2.Output()

	l.nor1.SetInputs(R, l.NQ)
    l.Q = l.nor1.Output()
	
}