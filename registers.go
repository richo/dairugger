package dairugger

import (
	"encoding/json"
)

type X64RegistersWrapper struct {
	Registers X64Registers `json: "registers"`
}

type X64Registers struct {
	RIP uint `json: "rip"`
	RAX uint `json: "rax"`
	RBX uint `json: "rbx"`
	RBP uint `json: "rbp"`
	RSP uint `json: "rsp"`
	RDI uint `json: "rdi"`
	RSI uint `json: "rsi"`
	RDX uint `json: "rdx"`
	RCX uint `json: "rdx"`
	R8  uint `json: "r8"`
	R9  uint `json: "r9"`
	R10 uint `json: "r10"`
	R11 uint `json: "r11"`
	R12 uint `json: "r12"`
	R13 uint `json: "r13"`
	R14 uint `json: "r14"`
	R15 uint `json: "r15"`
}

func (c *Client) GetX64Registers() (*X64Registers, error) {
	resp, err := c.Get("registers")

	if err != nil {
		return nil, err
	}

	var data X64RegistersWrapper
	err = json.Unmarshal(resp.Data, &data)

	return &data.Registers, nil
}
