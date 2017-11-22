package cdflib

import (
	"testing"
)

func TestFixedCondition(t *testing.T) {
	expected := 0.0
	actual := P(expected).P()
	if actual != expected {
		t.Errorf("P(x) != x. Expected %f, got %f", expected, actual)
	}	
	expected = 0.125
	actual = P(expected).P()
	if actual != expected {
		t.Errorf("P(x) != x. Expected %f, got %f", expected, actual)
	}
	expected = 1.0
	actual = P(expected).P()
	if actual != expected {
		t.Errorf("P(x) != x. Expected %f, got %f", expected, actual)
	}
	expected = 0.0
	actual = P(-0.0001).P()
	if actual != expected {
		t.Errorf("Expected P(x<0) = 0, got %f.", actual)
	}
	expected = 1.0
	actual = P(1.0001).P()
	if actual != expected {
		t.Errorf("Expected P(x>1) = 1, got %f.", actual)
	}
}
	
func TestNot(t *testing.T) {
	type Case struct {
		in Condition
		expected float64
	}
	cases := []Case{
		{True(), 0.0},
		{False(), 1.0},
	}
	for i, c := range(cases) {
		actual := Not(c.in).P()
		if actual != c.expected {
			t.Errorf("Not() case %d: expected %f, got %f.", i, c.expected, actual)
		}
	}
}

func TestAnd(t *testing.T) {
	type Case struct {
		conds []Condition
		expected float64
	}
	cases := []Case{
		{[]Condition{}, 1.0},
		{[]Condition{False()}, 0.0},
		{[]Condition{True()}, 1.0},
		{[]Condition{False(),False()}, 0.0},
		{[]Condition{False(),P(0.25)}, 0.0},
		{[]Condition{P(0.5),P(0.5)}, 0.25},
		{[]Condition{True(),P(0.5)}, 0.5},
		{[]Condition{True(),False()}, 0.0},
		{[]Condition{True(),True()}, 1.0},
		{[]Condition{True(),True(),True()}, 1.0},
		{[]Condition{True(),True(),False()}, 0.0},
	}
	for i, c := range(cases) {
		actual := And(c.conds...).P()
		if actual != c.expected {
			t.Errorf("And() case %d: expected %f, got %f.", i, c.expected, actual)
		}
	}
}

func TestOr(t *testing.T) {
	type Case struct {
		conds []Condition
		expected float64
	}
	cases := []Case{
		{[]Condition{}, 0.0},
		{[]Condition{False()}, 0.0},
		{[]Condition{True()}, 1.0},
		{[]Condition{False(),False()}, 0.0},
		{[]Condition{True(),False()}, 1.0},
		{[]Condition{True(),True()}, 1.0},
		{[]Condition{False(),P(0.25)}, 0.25},
		{[]Condition{P(0.5),P(0.5)}, 0.75},
		{[]Condition{True(),True(),True()}, 1.0},
		{[]Condition{True(),True(),False()}, 1.0},
		{[]Condition{False(),False(),False()}, 0.0},
	}
	for i, c := range(cases) {
		actual := Or(c.conds...).P()
		if actual != c.expected {
			t.Errorf("Or() case %d: expected %f, got %f.", i, c.expected, actual)
		}
	}
}
