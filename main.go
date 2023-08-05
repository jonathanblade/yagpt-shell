package main

import "github.com/jonathanblade/yagpt-shell/cmd"

var Version = "0.1.0-dev"

func main() {
	cmd.Execute(Version)
}
