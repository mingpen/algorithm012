package main

func main() {

}
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	for _, s := range strs {
		arr := getArr(s)
		m[arr] = append(m[arr], s)
	}
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func getArr(s string) (arr [26]byte) {
	cnt := len(s)
	for i := 0; i < cnt; i++ {
		arr[s[i]-'a']++
	}
	return
}
