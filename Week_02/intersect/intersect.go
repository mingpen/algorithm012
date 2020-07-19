package main

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	res := []int{}
	for _, v := range nums2 {
		cnt, ok := m[v]
		if !ok || cnt < 1 {
			continue
		}
		m[v]--
		res = append(res, v)
	}
	return res
}
