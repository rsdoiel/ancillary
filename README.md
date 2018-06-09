
# Ancillary

Like Electron but so much less.

Ancillary provides a simple way of creating a GUI leveraging a web browser
and Go's native http library.  Electron has become very polular for making
cross platform GUI apps because the UI can be implemented using
standard web technologies.  This is fine on computers like laptops and
desktops but a little more challanging for single board systems like
a Raspberry Pi.  Electron isn't the most slim of platforms. It
can be challenging to compile from scratch on a system where the 
dependencies are not readily available.  Ancillary seeks to be a poorman's
Electron providing only a simple web service over http on an available port
on localhost using one time tokens for keeping the http session private. Ancillary is built
with Go which provides enough in the way of standard libraries to remove
the dependency problem of moving Electron to a new platform. It should
work on an platform that Go can run on and will attempt to use whatever
the default web browser is available on the system. This means it can work
on a web browser already optimized for that platform.

## An Ancillary App

+ import the Ancillary package into your Go application
+ populate a map[string]byte[]{} with your assets
+ create an Ancillary structure with inky.CreateApp(so.Argv, map[string]byte[]{})
+ envoke app.Run() to start up your application
+ Ancillary will start a web service on the first available port (e.g. port 0), show the URL and try to start the default web browser pointing at the service.




