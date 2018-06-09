// Ancillary, a poorman's Web browser as GUI platform for Go apps.
// It similar to Electronic but so much less.
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
package ancillary

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

const (
	Version = `v0.0.0-idea`
)

// Ancillary is the struct holding the ancillary application environment
type Ancillary struct {
	Port       string            // The port number as string established by the launched service, default value is ":8111"
	WebBrowser string            // The path to the web browser for running the UI
	Htdocs     string            // Htdocs document root for assets, defaults to ""
	args       []string          // usually set to a copy of os.Argv
	assets     map[string][]byte // The assets (e.g. HTML, CSS, JavaScript) for your service
	secret     string            // The onetime key generate to establish a private channel of communications between the browser and Ancillary web service
}

// CreateApp creates and populates an Ancillary struct with
// defaults including web browser targetted, args, and assets
// (e.g. map to HTML5/CSS/JavaScript assets associated with the UI
// of the app).
func CreateApp(webBrowser string, args []string, assets map[string][]byte) *Ancillary {
	app := new(Ancillary)
	//FIXME: need to figure out how to determine a free port above 8000
	// It should be possible to run multiple Ancillary apps with our
	// collision of the port number.
	app.Port = ":8111"
	app.WebBrowser = webBrowser
	app.args = args
	app.assets = assets
	return app
}

// ResetAssets replaces the existing asset map with an empty one
func (i *Ancillary) ResetAssets() {
	i.assets = map[string][]byte{}
}

// SetAsset will add an asset to Ancillary struct
func (i *Ancillary) SetAsset(key string, value []byte) {
	i.assets[key] = value
}

// RunApp is responsible for creating a web service and launching
// the previously designated web browser, it takes two funcs
// the first is an initializaiton function, the second an HTTP HandleFunc
// which gets bound to the web service, the latter may be nil for
// this apps that don't need more service then loading assets.
func (i *Ancillary) RunApp(init func(*Ancillary) error, defaultHandler http.Handler) error {
	// Initialize the app
	if err := init(i); err != nil {
		return err
	}
	// Setup web service
	if len(i.assets) > 0 {
		for p, src := range i.assets {
			http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
				//FIXME: Manage session here.
				fmt.Fprintf(w, "%s", src)
			})
		}
	}
	// Token to maintain a private single user channel between
	// web service and browser
	if i.Htdocs != "" {
		http.Handle("/", http.FileServer(http.Dir(i.Htdocs)))
	}
	if defaultHandler != nil {
		http.Handle("/", defaultHandler)
	}
	// Find an available port and launch a web service in a go routine
	u := new(url.URL)
	u.Scheme = "http"
	u.Host = "localhost" + i.Port

	// Launch browser and connect to web service
	//i.args = append(i.args, "--new-instance")
	//i.args = append(i.args, "--safe-mode")
	i.args = append(i.args, "--new-window")
	//q := u.Query()
	//q.Set("app", i.secret)
	//u.RawQuery = q.Encode()
	i.args = append(i.args, u.String())

	if i.WebBrowser == "" {
		return fmt.Errorf("web browser not configured, try setting WEB_BROWSER variable in your shell")
	}
	go func() {
		time.Sleep(200 * time.Millisecond)
		log.Printf("Launching %s %s", i.WebBrowser, strings.Join(i.args, " "))
		// NOTE: If firefox isn't already running, closing the browser
		// should also shutdown the app.
		cmd := exec.Command(i.WebBrowser, i.args...)
		if err := cmd.Start(); err != nil {
			log.Printf("Browser: %s", err)
		}
		err := cmd.Wait()
		if err != nil {
			log.Printf("Finished, %s with error: %s", i.WebBrowser, err)
			os.Exit(1)
		} else {
			log.Printf("Finished: %s, OK", i.WebBrowser)
			os.Exit(0)
		}
	}()

	log.Printf("Listening on %s", u.String())
	err := http.ListenAndServe(u.Host, nil)
	if err != nil {
		log.Fatalf("Web Service: %s", err)
	}
	return nil
}
