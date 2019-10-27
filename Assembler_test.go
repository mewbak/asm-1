package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestAssembler(t *testing.T) {
	a := asm.New()
	a.Exit(0)
	assert.DeepEqual(t, a.Bytes(), []byte{0xb8, 0x3c, 0x00, 0x00, 0x00, 0xbf, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x05})
}

func TestMoveRegisterNumber(t *testing.T) {
	a := asm.New()
	a.MoveRegisterNumber("rax", int32(0xff))
	assert.DeepEqual(t, a.Bytes(), []byte{0xb8, 0xff, 0, 0, 0})
}

func TestMoveRegisterNumber2(t *testing.T) {
	a := asm.New()
	a.MoveRegisterNumber("rcx", int32(0xff))
	assert.DeepEqual(t, a.Bytes(), []byte{0xb8 + 1, 0xff, 0, 0, 0})
}

func TestMoveRegisterRegister(t *testing.T) {
	a := asm.New()
	a.MoveRegisterRegister("rax", "rax")
	assert.DeepEqual(t, a.Bytes(), []byte{0x48, 0x89, 0xc0})
}

func TestMoveRegisterRegister2(t *testing.T) {
	a := asm.New()
	a.MoveRegisterRegister("rcx", "r9")
	assert.DeepEqual(t, a.Bytes(), []byte{0x4c, 0x89, 0xc9})
}

func TestMoveRegisterRegister3(t *testing.T) {
	a := asm.New()
	a.MoveRegisterRegister("r9", "rcx")
	assert.DeepEqual(t, a.Bytes(), []byte{0x49, 0x89, 0xc9})
}

func TestPushRegister(t *testing.T) {
	a := asm.New()
	a.PushRegister("rbp")
	assert.DeepEqual(t, a.Bytes(), []byte{0x55})
}

func TestPushRegister2(t *testing.T) {
	a := asm.New()
	a.PushRegister("r9")
	assert.DeepEqual(t, a.Bytes(), []byte{0x41, 0x51})
}

func TestPopRegister(t *testing.T) {
	a := asm.New()
	a.PopRegister("rbp")
	assert.DeepEqual(t, a.Bytes(), []byte{0x5d})
}

func TestPopRegister2(t *testing.T) {
	a := asm.New()
	a.PopRegister("r9")
	assert.DeepEqual(t, a.Bytes(), []byte{0x41, 0x59})
}
