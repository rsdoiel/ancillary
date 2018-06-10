//
// ancillary is a proof of concept for a light weight binding of web service// and locally installed web browser for deliverying a UI for a Go based app.
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
package ancillary

//
// Quick sketch of the idea behind ancillary.
//
// ```
//      import (
//         "os"
//         "github.com/rsdoiel/ancillary"
//      )
//
//      func appInit(i *Ancillary) error {
//			appName:= path.Base(os.Argv[0])
//          log.Println("Nothing to initialize for %s\n", appName)
//      }
//
//      func HandleHelloWorld(next http.Handler) Handler {
//			return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
//              if r.URL.Path == "/" || r.URL.Path == "/helloworld" {
//                  w.Write("Hi there!")
//              }
//				next.ServeHTTP(w, r)
//			})
//      }
//
//      func main() {
//      	assets := loadAssets()
//      	app := ancillary.CreateApp("firefox", os.Args)
//      	msg, err := app.Run(appInit, HandleHelloWord)
//			if err != nil {
//          	log.Fatal(err)
//      	}
//          log.Println(msg)
//      }
// ```
//
