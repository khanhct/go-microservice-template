package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Context for "ls" command
type LsCommand struct {
	All bool
}

type flagConfig struct {
	All string
}

func (l *LsCommand) run(c *kingpin.ParseContext) error {
	fmt.Printf("all=%v\n", l.All)
	return nil
}

func configureLsCommand(app *kingpin.Application) {
	c := &LsCommand{}
	ls := app.Command("ls", "List files.").Action(c.run)
	ls.Flag("all", "List all files.").Short('a').BoolVar(&c.All)
}

func main() {
	app := kingpin.New("modular", "My modular application.")
	// configureLsCommand(app)
	cfg := flagConfig{
		All: "localhost",
	}
	app.Flag("listen-address", "Address to listen on for UI, API, and telemetry.").Default("0.0.0.0:9090").StringVar(&cfg.All)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
