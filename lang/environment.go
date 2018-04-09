package lang

import (
	"io"
)

type Environment interface {
	Out() io.Writer
	Err() io.Writer
	GetVar(name string) (Expression, bool)
	SetVar(name string, expr Expression)
}
