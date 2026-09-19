package helpers

import (
	"log"

	"github.com/chzyer/readline"
)

func CreateReader(prompt string) *readline.Instance {
	r, err :=  readline.NewEx(&readline.Config {
		Prompt: prompt,
	})

	if err != nil {
		log.Fatal(err)
	}

	return r
}
