package gf

import "log"

// Check err and logs/quits if there is
func Check(err error) { if err != nil { log.Fatal(err) } }
func Generate(text string) { text = "\n" + text; log.Fatal(text) }
func NonFatal(err error) { if err != nil { log.Print(err) } }
