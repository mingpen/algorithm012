package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make(map[rune]int)
	tmap := make(map[rune]int)

	gmap(smap, s)
	gmap(tmap, t)

	for k, v := range smap {
		if tmap[k] != v {
			return false
		}
	}
	return true
}

func gmap(m map[rune]int, s string) {
	for _, r := range s {
		m[r]++
	}
}
