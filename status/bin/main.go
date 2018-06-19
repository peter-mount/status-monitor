package main

import (
  "github.com/peter-mount/golib/kernel"
  "github.com/peter-mount/status-monitor/status"
  "log"
)

func main() {
  err := kernel.Launch( &status.Status{} )
  if err != nil {
    log.Fatal( err )
  }
}
