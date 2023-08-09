package main

import "github.com/jonathanblade/yagpt-shell/cmd"

var Version = "v0.0.0-dev"

func main() {
	cmd.Execute(Version)
}
