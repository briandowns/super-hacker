package templates

import (
	"errors"
	"math/rand"

	"github.com/briandowns/super-hacker/templates/c"
	"github.com/briandowns/super-hacker/templates/golang"
	"github.com/briandowns/super-hacker/templates/haskell"
	"github.com/briandowns/super-hacker/templates/python"
	"github.com/briandowns/super-hacker/templates/scala"
)

// goTemplates holds all active Go templates.
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

// pythonTemplates holds all active Python templates.
var pythonTemplates = []string{
	python.InterpreterPypy,
	python.InterpreterGenerator,
}

// scalaTemplates holds all active Scala templates.
var scalaTemplates = []string{
	scala.ScalaPhrases,
}

// cTemplates holds all active C templates.
var cTemplates = []string{
	c.BinTree,
	c.RedBlack,
	c.SimpleMap,
}

// haskellTemplates holds all active Haskell templates.
var haskellTemplates = []string{
	haskell.EtcPasswd,
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
	case "c":
		return cTemplates[rand.Intn(len(cTemplates))], nil
	case "haskell":
		return haskellTemplates[rand.Intn(len(haskellTemplates))], nil
	default:
		return "", errors.New("unsupported language")
	}
}
