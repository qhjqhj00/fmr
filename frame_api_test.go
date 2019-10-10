package fmr

import (
	"testing"
)

func TestMatchFrames(t *testing.T) {
	cases := []string{
		`两点半订黑暗森林`,
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
