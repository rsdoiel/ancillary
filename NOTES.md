
# Basic idea

+ load assets for a "web app" UI as a map[string][]byte
    + this allows the assets to the UI to be easily compiled in with the binary application
+ create an Ancillary struct using inky.CreateApp(os.Argv, map[string][]byte{}{}) 
    + We need a minimal amount of metadata to run our application environment, this should be help in the struct returned by CreateApp
+ "Run" the app by invoke Run() on the inky struct
    + takes a handler that maps your app's requests to your applications "server side" functionality (e.g open files, perform non-UI calculations, etc)
    + starts a web service on an available port (e.g. using port 8111)
    + attempts to launch a web browser pointing at the service
    + provides the applications' runtime

## Ancillary ideas

+ Ancillary package should only provide the binding of web service to web browser launch and maybe a web hook to end the web service, nothing else
+ Executing the app should reveal the localhost URL and any additional info to manually launch a web browser
+ We need some way to prevent another user on the system from sniffing the session, probably auto-generate one time passwords and tokens that are resolved on launch of the web browser's initial negotiation with the webpage
+ If no other web browser instance is running shutting down the web browser should also shutdown the app's web service

# Benefits

+ By relying only on Go's standard library an Ancillary App should be able to be cross compiled to all platforms where Go runs and a web browser is available+ Shipping an Ancillary App is was easy to copying a Go static binary
+ Ancillary should be reasonable light weight assuming the platform has an optimized web browser
+ An inky app provides the API "server side" that the app needs, no other resources get exposed (i.e. no default exposure is provided, no file system access, no process access)
+ An Ancillary UI is created using standard HTML5, CSS, JavaScript/WebASM that the browser supports so you could limit yourself to HTML5 and create UI that run in Lynx on an embedded system
+ GopherJS lets you write programming logic and compile to JavaScript today, Go's compiler is expected to compile to WebASM in an upcoming release, in theory an Ancillary App could be written completely in Go
+ Ancillary App is a Go package and thus could be exposed to other languages as a shared library, e.g. Python, Julia
+ You can use Go scripting implementations like Lau, JavaScript to provide customization in your app

