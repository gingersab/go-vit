package gulplog

import (
	"log"
	"os"
)

var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Error = log.New(os.Stdout, "\u001b[32mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Fatal = log.New(os.Stdout, "\u001b[31mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Panic = log.New(os.Stdout, "\u001b[30mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
