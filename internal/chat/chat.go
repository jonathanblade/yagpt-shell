package chat

import (
	"github.com/jonathanblade/yagpt-shell/internal/chat/ui"
	"github.com/jonathanblade/yagpt-shell/internal/config"
	"github.com/jonathanblade/yagpt-shell/internal/logger"
)

func Run(debug bool) {
	config := config.Read()
	logger := logger.New(debug)
	ui := ui.New(config, logger)
	ui.Run(debug)
}
