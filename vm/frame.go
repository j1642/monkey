package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int // instructional pointer
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
