// Copyright 2014 The Go Authors. All rights reserved.
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
				if i > 1 && consonants[s[i-1:i]] {
					fmt.Printf("っ")
				}
				wid = 3
				fmt.Printf("%s", cha)
				continue
			}
		}
		if i+2 <= len(s) {
			ka, ok := two[s[i:i+2]]
			if ok {
				if i > 1 && consonants[s[i-1:i]] {
					fmt.Printf("っ")
				}
				wid = 2
				fmt.Printf("%s", ka)
				continue
			}
		}
		wid = 1
		a, ok := one[s[i:i+1]]
		if ok {
			if i > 1 && consonants[s[i-1:i]] {
				fmt.Printf("っ")
			}
			fmt.Printf("%s", a)
			continue
		}
	}
	fmt.Print("\n")
}

var consonants = map[string]bool{
	"b": true,
	"c": true,
	"d": true,
	"f": true,
	"g": true,
	"h": true,
	"j": true,
	"k": true,
	"l": true,
	"p": true,
	"q": true,
	"r": true,
	"s": true,
	"t": true,
	"v": true,
	"w": true,
	"x": true,
	"y": true,
	"z": true,
}

// note the absense of n and m

var one = map[string]string{
	"a": "あ",
	"i": "い",
	"u": "う",
	"e": "え",
	"o": "お",
	"n": "ん",
}

var two = map[string]string{
	"xa": "ぁ",
	"xi": "ぃ",
	"xu": "ぅ",
	"xe": "ぇ",
	"xo": "ぉ",

	"ka": "か",
	"ki": "き",
	"ku": "く",
	"ke": "け",
	"ko": "こ",

	"ga": "が",
	"gi": "ぎ",
	"gu": "ぐ",
	"ge": "げ",
	"go": "ご",

	"sa": "さ",
	"su": "す",
	"se": "せ",
	"so": "そ",

	"za": "ざ",
	"ji": "じ",
	"zu": "ず",
	"ze": "ぜ",
	"zo": "ぞ",

	"ja": "じゃ",
	"ju": "じゅ",
	"jo": "じょ",

	"ta": "た",
	"te": "て",
	"to": "と",

	"ti": "てぃ",
	"tu": "とぅ",

	"di": "でぃ",
	"du": "どぅ",

	"da": "だ",
	"de": "で",
	"do": "ど",

	"na": "な",
	"ni": "に",
	"nu": "ぬ",
	"ne": "ね",
	"no": "の",

	"ha": "は",
	"hi": "ひ",
	"fu": "ふ",
	"he": "へ",
	"ho": "ほ",

	"ba": "ば",
	"bi": "び",
	"bu": "ぶ",
	"be": "べ",
	"bo": "ぼ",

	"pa": "ぱ",
	"pi": "ぴ",
	"pu": "ぷ",
	"pe": "ぺ",
	"po": "ぽ",

	"fa": "ふぁ",
	"fi": "ふぃ",
	"fe": "ふぇ",
	"fo": "ふぉ",

	"ma": "ま",
	"mi": "み",
	"mu": "む",
	"me": "め",
	"mo": "も",

	"ya": "や",
	"yu": "ゆ",
	"ye": "いぇ",
	"yo": "よ",

	"ra": "ら",
	"ri": "り",
	"ru": "る",
	"re": "れ",
	"ro": "ろ",

	"wa": "わ",
	//"wi": "ゐ",
	//"we": "ゑ",
	"wo": "を",

	//"wa": "うぁ",
	"wi": "うぃ",
	"we": "うぇ",
	//"wo": "うぉ",

	"va": "ゔぁ",
	"vi": "ゔぃ",
	"vu": "ゔ",
	"ve": "ゔぇ",
	"vo": "ゔぉ",
}

var three = map[string]string{
	"kya": "きゃ",
	"kyu": "きゅ",
	"kyo": "きょ",

	"gya": "ぎゃ",
	"gyu": "ぎゅ",
	"gyo": "ぎょ",

	"shi": "し",

	"sha": "しゃ",
	"shu": "しゅ",
	"sho": "しょ",

	"chi": "ち",
	"tsu": "つ",

	"dhi": "ぢ",
	"dhu": "づ",

	"cha": "ちゃ",
	"chu": "ちゅ",
	"che": "ちぇ",
	"cho": "ちょ",

	"dha": "ぢゃ",
	//"dhu": "ぢゅ",
	"dhe": "ぢぇ",
	"dho": "ぢょ",

	"nya": "にゃ",
	"nyu": "にゅ",
	"nyo": "にょ",

	"hya": "ひゃ",
	"hyu": "ひゅ",
	"hyo": "ひょ",

	"bya": "びゃ",
	"byu": "びゅ",
	"byo": "びょ",

	"pya": "ぴゃ",
	"pyu": "ぴゅ",
	"pyo": "ぴょ",

	"mya": "みゃ",
	"myu": "みゅ",
	"myo": "みょ",

	"xya": "ゃ",
	"xyu": "ゅ",
	"xyo": "ょ",

	"rya": "りゃ",
	"ryu": "りゅ",
	"ryo": "りょ",
}

var hiragana2romanji = map[string]string{
	"a": "あ",
	"i": "い",
	"u": "う",
	"e": "え",
	"o": "お",

	"xa": "ぁ",
	"xi": "ぃ",
	"xu": "ぅ",
	"xe": "ぇ",
	"xo": "ぉ",

	"ka": "か",
	"ki": "き",
	"ku": "く",
	"ke": "け",
	"ko": "こ",

	"ga": "が",
	"gi": "ぎ",
	"gu": "ぐ",
	"ge": "げ",
	"go": "ご",

	"kya": "きゃ",
	"kyu": "きゅ",
	"kyo": "きょ",

	"gya": "ぎゃ",
	"gyu": "ぎゅ",
	"gyo": "ぎょ",

	"sa":  "さ",
	"shi": "し",
	"su":  "す",
	"se":  "せ",
	"so":  "そ",

	"za": "ざ",
	"ji": "じ",
	"zu": "ず",
	"ze": "ぜ",
	"zo": "ぞ",

	"sha": "しゃ",
	"shu": "しゅ",
	"sho": "しょ",

	"ja": "じゃ",
	"ju": "じゅ",
	"jo": "じょ",

	"ta":  "た",
	"chi": "ち",
	"tsu": "つ",
	"te":  "て",
	"to":  "と",

	"ti": "てぃ",
	"tu": "とぅ",

	"di": "でぃ",
	"du": "どぅ",

	"da":  "だ",
	"dhi": "ぢ",
	"dhu": "づ",
	"de":  "で",
	"do":  "ど",

	"cha": "ちゃ",
	"chu": "ちゅ",
	"che": "ちぇ",
	"cho": "ちょ",

	"dha": "ぢゃ",
	//"dhu": "ぢゅ",
	"dhe": "ぢぇ",
	"dho": "ぢょ",

	"na": "な",
	"ni": "に",
	"nu": "ぬ",
	"ne": "ね",
	"no": "の",

	"nya": "にゃ",
	"nyu": "にゅ",
	"nyo": "にょ",

	"ha": "は",
	"hi": "ひ",
	"fu": "ふ",
	"he": "へ",
	"ho": "ほ",

	"hya": "ひゃ",
	"hyu": "ひゅ",
	"hyo": "ひょ",

	"ba": "ば",
	"bi": "び",
	"bu": "ぶ",
	"be": "べ",
	"bo": "ぼ",

	"bya": "びゃ",
	"byu": "びゅ",
	"byo": "びょ",

	"pa": "ぱ",
	"pi": "ぴ",
	"pu": "ぷ",
	"pe": "ぺ",
	"po": "ぽ",

	"pya": "ぴゃ",
	"pyu": "ぴゅ",
	"pyo": "ぴょ",

	"fa": "ふぁ",
	"fi": "ふぃ",
	"fe": "ふぇ",
	"fo": "ふぉ",

	"ma": "ま",
	"mi": "み",
	"mu": "む",
	"me": "め",
	"mo": "も",

	"mya": "みゃ",
	"myu": "みゅ",
	"myo": "みょ",

	"ya": "や",
	"yu": "ゆ",
	"ye": "いぇ",
	"yo": "よ",

	"xya": "ゃ",
	"xyu": "ゅ",
	"xyo": "ょ",

	"ra": "ら",
	"ri": "り",
	"ru": "る",
	"re": "れ",
	"ro": "ろ",

	"rya": "りゃ",
	"ryu": "りゅ",
	"ryo": "りょ",

	"wa": "わ",
	//"wi": "ゐ",
	//"we": "ゑ",
	"wo": "を",

	//"wa": "うぁ",
	"wi": "うぃ",
	"we": "うぇ",
	//"wo": "うぉ",

	"va": "ゔぁ",
	"vi": "ゔぃ",
	"vu": "ゔ",
	"ve": "ゔぇ",
	"vo": "ゔぉ",

	"n": "ん",
}
