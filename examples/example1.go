package main

import (
	"fmt"
	"os"

	"github.com/eidolon/console"
	"github.com/eidolon/console/parameters"
)

var name = "World"

func main() {
	application := console.NewApplication("eidolon/console", "0.1.0")
	application.AddCommand(console.Command{
		Name:        "greet",
		Description: "Greet's the given user, or the world.",
		Help:        "You don't have to specify a name.",
		Configure: func(definition *console.Definition) {
			definition.AddOption(
				parameters.NewStringValue(&name),
				"-n, --name[=NAME]",
				"Provide a name for the greeting.",
			)
		},
		Execute: func(input *console.Input, output *console.Output) error {
			fmt.Printf("Hello, %s!\n", name)
			return nil
		},
	})

	code := application.Run(os.Args[1:])

	os.Exit(code)
}
