package fmr

import (
	"testing"
)

func TestMatchFrames(t *testing.T) {
	cases := []string{
		`预定黑暗森林`,
	}
	g, err := GrammarFromFile("mr.grammar")
	t.Log(g)
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
