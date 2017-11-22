package cdflib

/*
An interface that represents a boolean-typed random variable.
Think of it as something that can either happen or not happen
with some probability.
*/
type Condition interface {
	/*
	  Returns the probability that the condition will be true.
	*/
	P() float64
}

type fixedCondition struct {
	p float64
}

func (f *fixedCondition) P() float64 {
	return f.p
}

/*
Returns a condition that always happens.
*/
func True() Condition {
	return &fixedCondition{1.0}
}

/*
Returns a condition that never happens.
*/
func False() Condition {
	return &fixedCondition{0.0}
}

/*
Creates a condition that happens with a known, fixed probability.
*/
func P(p float64) Condition {
	if p <= 0.0 {
		return False()
	}
	if p >= 1.0 {
		return True()
	}
	return &fixedCondition{p}
}

type not struct {
	cond Condition	
}

func (n *not) P() float64 {
	return 1.0 - n.cond.P()
}

/*
Logical negation of a condition.
*/
func Not(cond Condition) Condition {
	return &not{cond}
}

type and struct {
	conds []Condition	
}

func (a *and) P() float64 {
	p := 1.0
	for _, c := range(a.conds) {
		p *= c.P()
	}
	return p
}

/*
Logical AND operation on independent events. 
*/
func And(conds ...Condition) Condition {
	switch len(conds) {
	case 0:
		return True()
	case 1:
		return conds[0]
	default:
		return &and{conds}
	}
}

type or struct {
	conds []Condition	
}

func (o *or) P() float64 {
	p := 1.0
	for _, c := range(o.conds) {
		p *= 1.0 - c.P()
	}
	return 1.0 - p
}

/*
Logical OR operation on independent events. 
*/
func Or(conds ...Condition) Condition {
	switch len(conds) {
	case 0:
		return False()
	case 1:
		return conds[0]
	default:
		return &or{conds}
	}
}
