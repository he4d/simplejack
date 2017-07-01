simplejack - a simple level based logger for golang
=======================================

![alt text](https://pbs.twimg.com/media/CNb26N1UAAAAHTw.jpg "Simple Jack")

Usage
-----
Pass the minimum loglevel and an io.Writer where the logger will write to.. simple as that!

### Loglevels:
+ TRACE
+ DEBUG
+ INFO
+ WARNING
+ ERROR
+ FATAL

### Example message:
+ LOGLEVEL: 2017/07/01 12:00:00 example.go:18: Example logmessage..

Usage examples:
--------------

+ Write to os.Stdout

```golang
package main

import (
	"os"

	"github.com/he4d/simplejack"
)

func main() {
	logger := simplejack.New(simplejack.WARNING, os.Stdout)
	logger.Trace.Print("This will land in nirvana")
	logger.Warning.Print("This will print to stdout")
}
```

+ Write to a file

```golang
package main

import (
	"log"
	"os"

	"github.com/he4d/simplejack"
)

func main() {
	file, err := os.OpenFile("example.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file example.log : %s", err)
	}

	logger := simplejack.New(simplejack.DEBUG, file)
	logger.Trace.Print("This will land in nirvana")
	logger.Fatal.Fatal("This will write to the file example.log and os.Exit(1)")
}
```

References
----------

[simplejack was inspired by this blog post](https://www.goinggo.net/2013/11/using-log-package-in-go.html)


