//
// cmd/ancillary-demo/ancillary-demo.go is a demo app based on the Ancillary Go
// package. It can display a "Hello World" message or load a directory
// htdocs directory as an web service/browser app. Probably should morph
// into somemthing useful if Ancillary becomes more than a proof of concept.
//
// Copyright (c) 2018, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	// Main Package
	"github.com/rsdoiel/ancillary"
)

var (
	usage = `
USAGE: %s [OPTIONS]

Hello World in your web browser
`

	showHelp    bool
	showVersion bool
)

func main() {
	var err error

	appName := path.Base(os.Args[0])
	flag.BoolVar(&showHelp, "h", false, "Display help")
	flag.BoolVar(&showHelp, "help", false, "Display help")
	flag.BoolVar(&showVersion, "v", false, "Display version")
	flag.BoolVar(&showVersion, "version", false, "Display version")
	flag.Parse()

	if showHelp {
		fmt.Printf(usage, appName)
		flag.PrintDefaults()
		os.Exit(0)
	}

	if showVersion {
		fmt.Printf("%s\n", ancillary.Version)
		os.Exit(0)
	}

	browser := "firefox"
	browser, err = exec.LookPath(browser)
	if err != nil {
		browser = os.Getenv("WEB_BROWSER")
	}
	if browser == "" {
		log.Fatalf("Can't find Firefox or designated WEB_BROWSER")
	}
	// 1. Create an app
	log.Printf("Creating app %s", appName)
	app := ancillary.CreateApp(browser, flag.Args())

	// 2. Add any HTML and CSS assets
	log.Printf("Setting up %s", appName)
	for p, val := range HTML {
		if p == "index.html" {
			p = "README.html"
		}
		log.Printf("Adding asset %s", p)
		app.SetAsset(p, val)
	}
	// Add our CSS
	for p, val := range CSS {
		log.Printf("Adding asset %s", p)
		app.SetAsset(p, val)
	}
	// add our initial landing page.
	app.SetAsset("/", []byte(`
<!DOCTYPE html>
<html>
  <body>
  	<h1>Hello World!</h1>
	<p>This is a proof of concept demo of Ancillary</p>
	<ul>
		<li><a href="/time">Time</a> (dynamically calculated by service)</li>
		<li><a href="/helloworld">Hello World</a> (dynamically calculated by service)</li>
		<li><a href="README.html">README</a> (HTML page as asset)</li>
		<li><a href="NOTES.html">NOTES</a> (HTML page as asset)</li>
		<li><a href="license.html">LICENSE</a> (HTML page as asset)</li>
	</ul>
  </body>
</html>
`))
	// 3. RunApp passing in any additional initialization or http middleware service
	err = app.RunApp(func(app *ancillary.Ancillary) error {
		// This is the init function, we'd do any additional final
		// configuration before envoking the web service and launching web browser
		log.Printf("Launch web service and web browser for %s", appName)
		return nil
	}, func(next http.Handler) http.Handler {
		// This is an application specific handler adding
		// two behaviors via this middleware
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/time" {
				fmt.Fprintf(w, "%s", time.Now().String())
				return
			}
			if r.URL.Path == "/helloworld" {
				fmt.Fprintln(w, "Hi there!")
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	if err != nil {
		log.Fatal(err)
	}
}
