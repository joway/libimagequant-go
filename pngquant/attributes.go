package pngquant

import "errors"

/*
#include "libimagequant.h"
*/
import "C"

type Attributes struct {
	p *C.struct_liq_attr
}

// Callers MUST call Release() on the returned object to free memory.
func NewAttributes() (*Attributes, error) {
	pAttr := C.liq_attr_create()
	if pAttr == nil {
		return nil, errors.New("unsupported platform")
	}

	return &Attributes{p: pAttr}, nil
}

// Free memory. Callers must not use this object after Release has been called.
func (a *Attributes) Release() {
	C.liq_attr_destroy(a.p)
}
