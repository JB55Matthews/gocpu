A MIPS CPU Emulated in GO
---------------------------
A processor for MIPS emulated in Go. All components reduce down to simple gate structs, using no internal logic except gate connections. This makes this a full "emulation", and not simulated processor.
For example, while many simulated processors implement a register file as simply an array, here, we implement each register as 32 D flip-flops in parallel, which in turn are made of D latches, which is made of SR latches, which is made from gates. In this, we fully emulate the behaviour of a processor.

This is in early stages of being developed, so this is still very bare.
