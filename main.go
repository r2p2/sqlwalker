package main

import (
    "flag"
    "log"
)

var flagDebug bool
var flagDryRun bool
var flagSourceVersion int
var flagTargetVersion int

func init() {
    flag.BoolVar(&flagDebug, "d", false, "Enables the output of debug information.")
    flag.BoolVar(&flagDryRun, "dry", false, "If submitted, the goxdb tool will wrap a database transaction around the change set.")
    flag.IntVar(&flagSourceVersion, "src", 0, "Describes the source version from where the change set will start.")
    flag.IntVar(&flagTargetVersion, "tgt", 0, "Describes the target version from where the change set will start.")

    flag.Parse()
}

func main() {
    if flagDebug {
        log.Println("parameter flagDebug: ", flagDebug)
        log.Println("parameter flagDryRun: ", flagDryRun)
        log.Println("parameter flagSourceVersion: ", flagSourceVersion)
        log.Println("parameter flagTargetVersion: ", flagTargetVersion)
    }
}
