// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wasm

import (
	"fmt"
	// "log"
)

type Logger interface {
	Printf(string, ...interface{})
	Println(string, ...interface{})
}

var logger Logger

func init() {
	logger = NoopLogger{}
}

func SetLogger(l Logger) {
	logger = l
}

type NoopLogger struct{}

func (l NoopLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func (l NoopLogger) Println(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v...))
}

/*
import (
	"io/ioutil"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	SetDebugMode(false)
}

func SetDebugMode(dbg bool) {
	w := ioutil.Discard
	if dbg {
		w = os.Stderr
	}
	logger = log.New(w, "", log.Lshortfile)
}
*/
