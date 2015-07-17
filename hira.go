// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Hira accepts, as argument text or as standard input, romaji
text. It copies the text to standard output, converting
the text that matches the romaji format to hiragana.
Note that the output may not be accurate because of
false matches and the inability to generate things like
the tsu consonant-extending symbol.
*/
package main // import "robpike.io/cmd/hira"

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"robpike.io/nihongo"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, nihongo.HiraganaReader(os.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	text := nihongo.HiraganaString(strings.Join(os.Args[1:], " "))
	fmt.Printf("%s\n", text)
}
