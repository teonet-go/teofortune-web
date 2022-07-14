// Copyright 2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Teonet fortune web-server microservice. This is simple Teonet web-server
// micriservice application which get fortune message from Teonet Fortune
// microservice and show it in the site web page.
package main

import (
	"flag"
	"log"
	"time"

	"github.com/teonet-go/teonet"
)

const (
	appShort   = "teofortune-web"
	appName    = "Teonet fortune web-server microservice application"
	appLong    = ""
	appVersion = "0.0.3"
)

var appStartTime = time.Now()
var domain, fortune, monitor string

// Params is teonet command line parameters
var Params struct {
	appShort    string
	port        int
	httpAddr	string
	stat        bool
	hotkey      bool
	showPrivate bool
	loglevel    string
	logfilter   string
}

func main() {

	// Application logo
	teonet.Logo(appName, appVersion)

	// Parse application command line parameters
	flag.StringVar(&Params.appShort, "name", appShort, "application short name")
	flag.IntVar(&Params.port, "p", 0, "local port")
	flag.StringVar(&Params.httpAddr, "addr", "localhost:8088", "http server local address")
	flag.BoolVar(&Params.stat, "stat", false, "show statistic")
	flag.BoolVar(&Params.hotkey, "hotkey", false, "start hotkey menu")
	flag.BoolVar(&Params.showPrivate, "show-private", false, "show private key")
	flag.StringVar(&Params.loglevel, "loglevel", "NONE", "log level")
	flag.StringVar(&Params.logfilter, "logfilter", "", "log filter")
	//
	flag.StringVar(&domain, "domain", "", "domain name to process HTTP/s server")
	flag.StringVar(&fortune, "fortune", "", "fortune microservice address")
	flag.StringVar(&monitor, "monitor", "", "monitor address")
	//
	flag.Parse()

	// Check requered parameters
	teonet.CheckRequeredParams("fortune")

	// Initialize and run Teonet
	teo, err := newTeonet()
	if err != nil {
		log.Panic(err)
		return
	}

	// Initialize and run web-server
	err = newServe(domain, Params.httpAddr, teo)
	if err != nil {
		log.Panic(err)
		return
	}
}
