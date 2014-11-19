// +build cortexm3 cortexm4 cortexm4f

__attribute__ ((naked)) static
void runtime$noos$svcHandler() {
	asm volatile (
		"tst	lr, 4\n\t"
		"ite 	eq\n\t"
		"mrseq 	r0, msp\n\t"
		"mrsne	r0, psp\n\t"
		"b		runtime$noos$sv"
		:: "X" (runtime$noos$sv)
	);
}

static
void (*runtime$noos$p2f(uintptr p))() {
	union {uintptr in; void (*out)();} cast;
	cast.in = p;
	return cast.out;
}
