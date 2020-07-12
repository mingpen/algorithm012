package main

func main() {

}

//
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int) // value 与index的对应值
	for i, v := range nums {
		if i == 0 {
			m[v] = i
			continue
		}
		want := target - v
		j, ok := m[want]
		if ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

// 暴力求解
func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for j, w := range nums[i+1:] {
			if v+w == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}
