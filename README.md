simplejack - a simple logger for golang
=======================================

![alt text](https://pbs.twimg.com/media/CNb26N1UAAAAHTw.jpg "Simple Jack")

Just wrapping around the stdlib logger.
So simple..

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
	logger.Warning.Print("This will print out to stdout")
	//Result:
	//WARNING: 2017/07/01 09:29:44 example_stdout.go:12: This will print out to stdout
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
	logger.Fatal.Fatalf("This will write to the file example.log and os.Exit(1)")

	//Result:
	//FATAL: 2017/07/01 09:32:46 example_file.go:18: This will write to the file example.log and os.Exit(1)
}
```

References
----------

[simplejack was inspired by this blog post](https://www.goinggo.net/2013/11/using-log-package-in-go.html)


