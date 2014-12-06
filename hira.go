// Copyright 2014 Rob Pike. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Hira accepts as argument text romaji
text. It copies the text to standard output, converting
the text that matches the romaji format to hiragana.
Note that the output may not be accurate because of
false matches and the inability to generate things like
the tsu consonant-extending symbol.
*/
package main // import "robpike.io/cmd/hira"

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	do(strings.Join(os.Args[1:], " "))
}

func do(s string) {
	wid := 0
	for i := 0; i < len(s); i += wid {
		if i+3 <= len(s) {
			cha, ok := three[s[i:i+3]]
			if ok {
				wid = 3
				fmt.Printf("%s", cha)
				continue
			}
		}
		if i+2 <= len(s) {
			ka, ok := two[s[i:i+2]]
			if ok {
				wid = 2
				fmt.Printf("%s", ka)
				continue
			}
		}
		wid = 1
		a, ok := one[s[i:i+1]]
		if ok {
			fmt.Printf("%s", a)
			continue
		}
		fmt.Printf("%s", s[i:i+1])
	}
	fmt.Print("\n")
}

var one = map[string]string{
	"a": "あ",
	"i": "い",
	"u": "う",
	"e": "え",
	"o": "お",
	"n": "ん",
}

var two = map[string]string{
	"ka": "か",
	"ga": "が",
	"ki": "き",
	"gi": "ぎ",
	"ku": "く",
	"gu": "ぐ",
	"ke": "け",
	"ge": "げ",
	"ko": "こ",
	"go": "ご",
	"sa": "さ",
	"za": "ざ",
	"si": "し",
	"zi": "じ",
	"su": "す",
	"zu": "ず",
	"se": "せ",
	"ze": "ぜ",
	"so": "そ",
	"zo": "ぞ",
	"ta": "た",
	"da": "だ",
	"ti": "ち",
	"di": "ぢ",
	"tu": "つ",
	"du": "づ",
	"te": "て",
	"de": "で",
	"to": "と",
	"do": "ど",
	"na": "な",
	"ni": "に",
	"nu": "ぬ",
	"ne": "ね",
	"no": "の",
	"ha": "は",
	"va": "ば",
	"pa": "ぱ",
	"hi": "ひ",
	"vi": "び",
	"pi": "ぴ",
	"fu": "ふ",
	"bu": "ぶ",
	"pu": "ぷ",
	"he": "へ",
	"ve": "べ",
	"pe": "ぺ",
	"ho": "ほ",
	"vo": "ぼ",
	"po": "ぽ",
	"ma": "ま",
	"mi": "み",
	"mu": "む",
	"me": "め",
	"mo": "も",
	"ya": "や",
	"yu": "ゆ",
	"yo": "よ",
	"ra": "ら",
	"ri": "り",
	"ru": "る",
	"re": "れ",
	"ro": "ろ",
	"wa": "わ",
	"wi": "ゐ",
	"we": "ゑ",
	"wo": "を",
	"vu": "ゔ",
}

var three = map[string]string{
	"shi": "し",
	"chi": "ち",
	"tsu": "つ",
	"cha": "ちや", // TODO IS THIS RIGHT
	"chu": "ちゆ", // TODO IS THIS RIGHT
	"cho": "ちよ", // TODO IS THIS RIGHT
	// NEED MORE
}
