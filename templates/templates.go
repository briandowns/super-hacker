package templates

import (
	"errors"
	"math/rand"

	"github.com/briandowns/super-hacker/templates/golang"
	"github.com/briandowns/super-hacker/templates/python"
	"github.com/briandowns/super-hacker/templates/scala"
)

var goTemplates = []string{
	golang.SyscallSyscallUnix,
	golang.BytesReader,
	golang.ContainerRing,
	golang.GoAstFilter,
	golang.GoTokenPosition,
	golang.BufioBufio,
	golang.CryptoAESGCM,
	golang.OsStatUnix,
	golang.OsPath,
	golang.ErrorsWrap,
	golang.FmtFormat,
}

var pythonTemplates = []string{
	python.InterpreterPypy,
	python.InterpreterGenerator,
}

var scalaTemplates = []string{
	scala.ScalaPhrases,
}

// RandomGo selects a template from the template slice
// at random and returns it to the caller.
func Random(lang string) (string, error) {
	switch lang {
	case "go":
		return goTemplates[rand.Intn(len(goTemplates))], nil
	case "python":
		return pythonTemplates[rand.Intn(len(pythonTemplates))], nil
	case "scala":
		return scalaTemplates[rand.Intn(len(scalaTemplates))], nil
	default:
		return "", errors.New("unsupported language")
	}
}
