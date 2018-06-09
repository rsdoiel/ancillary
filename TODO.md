
# TODO

## Bugs

+ If browser is open already, it needs to be detected so os.Exit(0) is called auto-magically because PID changes after new page it handed off to existin browser instance

## Next

+ Use a default port range, finding the next avail port for the web service
+ Secure the communication channel between envoking Ancillary instance and web browser

## Someday, maybe

+ Consider adding a webhook in the web service that will close the web service from the web browser, this would allow us to launch in browser only apps and not tie up a port number (e.g. http://localhost:8111?exit=0, would close the web service leaving the browser running)
+ Make map[string][]byte a type of Assets
    + (a *Assets)CreateAssets()
    + (a *Assets)ResetAssets()
    + (a *Assets)SetAsset()
    + This could allow an appliction to manage assets outside the binary
    + This could let you load assets as sets, e.g. multi-language support
