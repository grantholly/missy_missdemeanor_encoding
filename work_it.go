package main

import(
	"fmt"
	"strings"
	"bytes"
	"os"
)

func isItWorth (it bool) bool {
	if it {
		return true
	}
	return false
}

func flipIt (s string) string {
	upsidedown := make(map[rune]string)
	
	upsidedown['a'] = "ɐ"
	upsidedown['b'] = "p"
	upsidedown['c'] = "c"
	upsidedown['d'] = "p"
	upsidedown['e'] = "ǝ"
	upsidedown['f'] = "ɟ"
	upsidedown['g'] = "ƃ"
	upsidedown['h'] = "ɥ"
	upsidedown['i'] = "ᴉ"
	upsidedown['j'] = "ɾ"
	upsidedown['k'] = "ʞ"
	upsidedown['l'] = "l"
	upsidedown['m'] = "ɯ"
	upsidedown['n'] = "u"
	upsidedown['o'] = "o"
	upsidedown['p'] = "d"
	upsidedown['q'] = "b"
	upsidedown['r'] = "ɹ"
	upsidedown['s'] = "s"
	upsidedown['t'] = "ʇ"
	upsidedown['u'] = "n"
	upsidedown['v'] = "ʌ"
	upsidedown['w'] = "ʍ"
	upsidedown['x'] = "x"
	upsidedown['y'] = "ʎ"
	upsidedown['z'] = "z"

	var flipped bytes.Buffer

	for _, c := range s {
		if _, ok := upsidedown[c]; ok {
			flipped.WriteString(upsidedown[c])
		} else {
			continue
		}
	}
	return flipped.String()
}

func putMyThingDown (thing string) string {
	return strings.ToLower(thing)
}

func reverseIt (s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func letMeWorkIt (s string) string {
	missy := putMyThingDown(s)
	fmt.Println(missy)
	missy = flipIt(missy)
	fmt.Println(missy)
	missy = reverseIt(missy)
	fmt.Println(missy)
	return missy
}

func main () {

	var args bytes.Buffer

	for _, arg := range os.Args[1:] {
		args.WriteString(arg)
	}

	// https://en.wikipedia.org/wiki/Work_It_(Missy_Elliott_song)
	letsDoThis := args.String()

	it := true

	if isItWorth(it) {
		letMeWorkIt(letsDoThis)
	}
}
