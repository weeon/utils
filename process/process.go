package process

import (
	"os"
	"os/signal"
	"syscall"

	"log/slog"
)

var (
	exitFuncs = []func(){}
)

func AddExitFuncs(fns ...func()) {
	exitFuncs = append(exitFuncs, fns...)
}

func WaitSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	slog.Info("recv signal ", "sig", sig)
	for _, fn := range exitFuncs {
		fn()
	}
}
