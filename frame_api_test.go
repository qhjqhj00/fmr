package fmr

import (
	"testing"
)

func TestMatchFrames(t *testing.T) {
	cases := []string{
		`两点半订黑暗森林`,
		`四点二十分订一下微笑改需求`,
		`两点半到三点四十`,
	}
	g, err := GrammarFromFile("mr.grammar")
	if err != nil {
		t.Error(err)
	}
	for _, c := range cases {
		fmrs, err := g.FrameFMR(c)
		if err != nil {
			t.Error(err)
		}
		t.Log(c, fmrs)
	}
}
