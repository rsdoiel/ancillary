//
//      import (
//         "os"
//         "github.com/rsdoiel/ancillary"
//      )
//
//      func loadAssets() map[string][]byte {
//          // NOTE: This is where you'd load any web UI content
//          // Ancillary expectes a map pointing at arrays of bytes
//          return map[string][]byte{}{
//          "/index.html": []byte(```
//          <!DOCTYPE html>
//          <html>
//            <body>
//            Hello World!
//            </body>
//          </html>
//          ```),
//			}
//      }
//
//      func appInit(i *Ancillary) error {
//			appName:= path.Base(os.Argv[0])
//          log.Println("Nothing to initialize for %s\n", appName)
//      }
//
//      func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
//           w.Write("Hi there!")
//      }
//
//      func main() {
//      	assets := loadAssets()
//      	app := ancillary.CreateApp(os.Args, assets)
//      	msg, err := app.Run(appInit, helloWorldHandler)
//			if err != nil {
//          	log.Fatal(err)
//      	}
//          log.Println(msg)
//      }
//
