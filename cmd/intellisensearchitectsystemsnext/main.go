// cmd/intellisensearchitectsystemsnext/main.go
package main

import (
"flag"
"log"
"os"

"intellisensearchitectsystemsnext/internal/intellisensearchitectsystemsnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := intellisensearchitectsystemsnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
