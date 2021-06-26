package ggdebug

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/loov/hrtime"
)

var output *os.File
var currentSession string
var start time.Duration

func BeginSession(name, path string) {
	if currentSession != "" {
		panic("Session already exists")
	}

	var err error
	output, err = os.Create(path)
	if err != nil {
		panic(err)
	}

	currentSession = name
	writeHeader()

	start = hrtime.Now()
}

func EndSession() {
	writeFooter()
	if err := output.Close(); err != nil {
		panic(err)
	}
	currentSession = ""
}

func writeHeader() {
	output.WriteString("{\"otherData\": {},\"traceEvents\":[{}")
}

func writeFooter() {
	output.WriteString("]}")
}

func writeProfile(name string, startTime time.Duration, duration time.Duration) {
	json := fmt.Sprintf(`,
	{
		"cat": "function",
		"dur": %d,
		"name": "%v",
		"ph": "X",
		"pid": 0,
		"tid": 0,
		"ts": %d
	}
	`, duration.Microseconds(), name, startTime.Microseconds()-start.Microseconds())

	output.WriteString(json)
}

type Timer struct {
	name      string
	startTime time.Duration
}

func Start() *Timer {
	fnName := "<unknown>"

	pc, _, _, ok := runtime.Caller(1)
	if ok {
		split := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		fnName = split[len(split)-1]
	}

	timer := Timer{
		name:      fnName,
		startTime: hrtime.Now(),
	}

	return &timer
}

func Stop(timer *Timer) {
	duration := hrtime.Since(timer.startTime)
	writeProfile(timer.name, timer.startTime, duration)
}
