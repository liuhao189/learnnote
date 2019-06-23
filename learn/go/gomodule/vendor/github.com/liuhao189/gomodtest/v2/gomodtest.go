package gomodtest

import (
	"errors"
	"fmt"
)

func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi,%s!", name), nil
	case "pr":
		return fmt.Sprintf("Oi,%s!", name), nil
	case "es":
		return fmt.Sprintf("!Hola,%s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour,%s!", name), nil
	default:
		return "", errors.New("unknown language " + lang)
	}
}
