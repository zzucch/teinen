package main

import (
	"github.com/zzucch/teinen/internal/anki"
	"github.com/zzucch/teinen/internal/read"
	"github.com/zzucch/teinen/internal/waitlog"
)

func main() {
  _, err := read.Read()
  if err != nil {
    waitlog.Fatal(err)
  }

	anki.Run()
}
