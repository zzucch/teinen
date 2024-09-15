package main

import (
	"github.com/zzucch/teinen/internal/anki"
	"github.com/zzucch/teinen/internal/parse"
	"github.com/zzucch/teinen/internal/read"
	"github.com/zzucch/teinen/internal/waitlog"
)

func main() {
	lines, err := read.Read()
	if err != nil {
		waitlog.Fatal(err)
	}

  result := parse.Parse(lines)

  for _, res := range result {
    println("word: "+res.Word)
    println("meaning: "+res.Meaning)
    println("info: "+res.Info)
  }

	anki.Run()
}
