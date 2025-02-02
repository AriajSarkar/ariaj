package main

import (
    "github.com/AriajSarkar/ariaj/cmd"
    "os"
)

// These variables are set during build by goreleaser
var (
    version = "dev"
    commit  = "none"
    date    = "unknown"
)

func main() {
    cmd.SetVersion(version, commit, date)
    if err := cmd.Execute(); err != nil {
        os.Exit(1)
    }
}
