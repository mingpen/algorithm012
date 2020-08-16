package main

func main() {

}

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 1
	}
	f1, f2, f3 := 0, 0, 0
	if s[l-1] != '0' {
		f1 = 1
	}
	if l == 1 {
		return f1
	}
	if s[l-2] == '0' {
		f2 = 0
	} else if s[l-2:] < "27" {
		f2 = 1 + f1
	} else {
		f2 = f1
	}
	f3, f2 = f2, f1
	for i := 3; i <= l; i++ {
		f1, f2 = f2, f3
		if s[l-i] == '0' {
			f3 = 0
			continue
		}
		t := 0
		if s[l-i:l-i+2] < "27" {
			t = f1
		}
		f3 = f2 + t
	}
	return f3
}
