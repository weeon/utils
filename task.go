package utils

import (
	"time"

	"github.com/weeon/contract"
	"github.com/weeon/log"
	"go.uber.org/zap/zapcore"
)

type Task struct {
	name string
	t    time.Duration
	fn   func() error

	Logger contract.Logger
}

type TaskOpt func(t *Task)

func NewTask(name string, t time.Duration, fn func() error, opts ...TaskOpt) *Task {
	logger, _ := log.NewLogger("/dev/null", zapcore.DebugLevel)
	task := &Task{
		name:   name,
		t:      t,
		fn:     fn,
		Logger: logger,
	}
	for _, opt := range opts {
		opt(task)
	}
	return task
}

func (t *Task) Run() {
	for {
		start := time.Now()
		err := t.fn()
		t.Logger.Infof("run task %s cost: %v error: %v  ", t.name, time.Since(start), err)
		time.Sleep(t.t)
	}
}
