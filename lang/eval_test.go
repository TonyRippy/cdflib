package lang

import (
	"testing"
)

// A helper function that will fail if the code fails to evaluate.
// Does not check the result.
func checkEval(t *testing.T, code string) {
	t.Helper()
	i := MakeInterpreter()
	_, err := i.Eval(code)
	if err != nil {
		t.Errorf("Unable to evaluate \"%s\": %s", code, err)
	}
}

// A helper function that will fail if the code evaluates.
// (Expects failure.)
func checkNotEval(t *testing.T, code string) {
	t.Helper()
	i := MakeInterpreter()
	_, err := i.Eval(code)
	if err == nil {
		t.Errorf("Expected evaluation of \"%s\" to fail.", code)
	}
}

// Checks that the code evaluates to an expected result. 
func checkEvalText(t *testing.T, code string, expected string) {
	t.Helper()
	i := MakeInterpreter()
	data, err := i.Eval(code)
	if (err != nil) {
		t.Errorf("Unable to evaluate \"%s\": %s", code, err)
		return
	}
	actual := data.String()
	if (expected != actual) {
		t.Errorf("When evaluating \"%s\": expected \"%s\", got \"%s\".", code, expected, actual)
	}
}

func TestEmpty(t *testing.T) {
	checkEvalText(t, "", "nil")
	checkEvalText(t, " ", "nil")
	checkEvalText(t, " \t\n ", "nil")
	checkEvalText(t, ";", "nil")
	checkEvalText(t, " ; ", "nil")
	checkEvalText(t, ";; ;;;", "nil")
}

func TestNumber(t *testing.T) {
	checkEvalText(t, "42", "42")
	checkEvalText(t, "42;", "42")
	checkEvalText(t, " 42  ", "42")
	checkEvalText(t, "42.0", "42")
	checkEvalText(t, "1.2345", "1.2345")
}

func TestAssignment(t *testing.T) {
	checkEvalText(t, "x = 3;\ny = 4; x", "3")
}

func TestAdd(t *testing.T) {
	checkEvalText(t, "1 + 1", "2")
	checkEvalText(t, "1+2+3", "6")
}

func TestMathPrecedence(t *testing.T) {
	checkEvalText(t, "5+4-3+2", "8")
	checkEvalText(t, "9-5", "4")
	checkEvalText(t, "5+4*3-2", "15")
	checkEvalText(t, "5+4-3*2", "3")
}

func TestApply(t *testing.T) {
	checkEval(t, "N(0,1)")
	checkNotEval(t, "N(0)") // No overload takes 1 argument.
	checkNotEval(t, "N(0,N(0,1))") // No overload takes a CDF as an argument.
}
