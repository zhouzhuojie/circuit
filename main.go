package main

import (
	"github.com/codeamp/circuit/cmd"
	_ "github.com/codeamp/circuit/plugins/codeamp"
	_ "github.com/codeamp/circuit/plugins/dockerbuilder"
	_ "github.com/codeamp/circuit/plugins/githubstatus"
	_ "github.com/codeamp/circuit/plugins/gitsync"
	_ "github.com/codeamp/circuit/plugins/heartbeat"
	_ "github.com/codeamp/circuit/plugins/kubernetes"
	_ "github.com/codeamp/circuit/plugins/route53"
	_ "github.com/codeamp/circuit/plugins/slack"
	_ "github.com/codeamp/circuit/plugins/smartprofiles"
)

func main() {
	cmd.Execute()
}
