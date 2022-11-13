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
	"os"
	"time"

	"github.com/teonet-go/teonet"
)

const (
	appShort = "teofortune-web"
	appName  = "Teonet fortune web-server microservice application"
	appLong  = `
		This is simple <a href="https://github.com/teonet-go">Teonet</a> 
		web-server microservice application which get fortune message from 
		<a href="https://github.com/teonet-go/teofortune">Teonet Fortune</a> 
		microservice and show it in the web page.<br>
		See source code at <a href="https://github.com/teonet-go/teofortune-web">
		https://github.com/teonet-go/teofortune-web</a>
	`
	appVersion = "0.0.5"

	appPort = "8080"
)

var appStartTime = time.Now()
var domain, fortune, monitor string

// Params is teonet command line parameters
var Params struct {
	appShort           string
	port               int
	httpAddr           string
	stat               bool
	hotkey             bool
	showPrivate        bool
	loglevel           string
	logfilter          string
	directConnectDelay int
}

func main() {

	// Application logo
	teonet.Logo(appName, appVersion)

	// Get HTTP port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = appPort
	}

	// Parse application command line parameters
	flag.StringVar(&Params.appShort, "name", appShort, "application short name")
	flag.IntVar(&Params.port, "p", 0, "local port")
	flag.StringVar(&Params.httpAddr, "addr", ":"+port, "http server local address")
	flag.BoolVar(&Params.stat, "stat", false, "show statistic")
	flag.BoolVar(&Params.hotkey, "hotkey", false, "start hotkey menu")
	flag.BoolVar(&Params.showPrivate, "show-private", false, "show private key")
	flag.StringVar(&Params.loglevel, "loglevel", "debugv", "log level")
	flag.StringVar(&Params.logfilter, "logfilter", "", "log filter")
	flag.IntVar(&Params.directConnectDelay, "directconnect", 0, "use 'direct connect' to pear delay")
	//
	flag.StringVar(&domain, "domain", "", "domain name to process HTTP/s server")
	flag.StringVar(&fortune, "fortune", "", "fortune microservice address")
	flag.StringVar(&monitor, "monitor", "", "monitor address")
	//
	flag.Parse()

	// Get fortune address from environment variable
	if len(fortune) == 0 {
		fortune = os.Getenv("TEO_FORTUNE")
	}

	// Check requered parameters
	teonet.CheckRequeredParams("fortune")

	// Initialize and run Teonet
	teo, err := newTeonet()
	if err != nil {
		log.Panic(err)
		return
	}

	// Initialize and run web-server
	err = newServe(domain, appLong, Params.httpAddr, teo)
	if err != nil {
		log.Panic(err)
		return
	}
}
