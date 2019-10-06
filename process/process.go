package process

import (
	"github.com/weeon/log"
	"os"
	"os/signal"
	"syscall"
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
	log.Infof("recv signal %v", sig)
	for _, fn := range exitFuncs {
		fn()
	}
}
