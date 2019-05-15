package utils

import (
	"fmt"
	"runtime/debug"
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

func SetTaskLogger(l contract.Logger) TaskOpt {
	return func(t *Task) {
		t.Logger = l
	}
}

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

func NewTaskAndRun(name string, t time.Duration, fn func() error, opts ...TaskOpt) {
	task := NewTask(name, t, fn, opts...)
	go task.Run()
}

func (t *Task) run() {
	start := time.Now()
	err := t.fn()
	var errMsg string
	if err != nil {
		errMsg = fmt.Sprintf("error %s", err.Error())
	}
	t.Logger.Infof("run task %s cost: %v error: %s  ", t.name, time.Since(start), errMsg)
	time.Sleep(t.t)
}

func (t *Task) Run() {
	timer := time.NewTimer(t.t)

	defer func() {
		if err := recover(); err != nil {
			t.Logger.Errorf("task %s crash error: %v stack:\n%v", t.name, err, debug.Stack())
		}
	}()

	t.run()

	for {
		select {
		case <-timer.C:
			t.run()
			timer.Reset(t.t)
		}
	}
}
