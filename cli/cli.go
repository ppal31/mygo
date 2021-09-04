package cli

import (
	"github.com/ppal31/mygo/cli/server"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

// application name
var application = "policy-mgmt"

// application description
var description = "description goes here" // TODO edit this application description

// Command parses the command line arguments and then executes a
// subcommand program.
func Command() {
	app := kingpin.New(application, description)
	server.Register(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
