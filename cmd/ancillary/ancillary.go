package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

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
	browser     string
	appPath     string

	assets = map[string][]byte{
		"/": []byte("Hello World!"),
	}
)

func main() {
	var err error

	appName := path.Base(os.Args[0])
	flag.BoolVar(&showHelp, "h", false, "Display help")
	flag.BoolVar(&showHelp, "help", false, "Display help")
	flag.BoolVar(&showVersion, "v", false, "Display version")
	flag.BoolVar(&showVersion, "version", false, "Display version")
	flag.StringVar(&browser, "browser", "firefox", "Set web browser to use")
	flag.StringVar(&appPath, "app", "", "Run a directory")
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

	assets := map[string][]byte{
		"/": []byte(`
<!DOCTYPE html>
<html>
  <body>
  	<h1>Hello World!</h1>
  </body>
</html>
`),
	}
	browser, err = exec.LookPath(browser)
	if err != nil {
		browser = os.Getenv("WEB_BROWSER")
	}
	if browser == "" {
		log.Fatalf("Can't find Firefox or designated WEB_BROWSER")
	}
	app := ancillary.CreateApp(browser, flag.Args(), assets)
	if appPath != "" {
		app.Htdocs = appPath
		app.ResetAssets()
	}

	err = app.RunApp(func(i *ancillary.Ancillary) error {
		log.Printf("Setting up %s", appName)
		return nil
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
