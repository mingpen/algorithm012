func validMountainArray(A []int) bool {
cnt := len(A) 
if cnt < 3 || A[cnt-1] >= A[cnt-2] {
return false
}
// 状态 先上坡，再下坡
up := true
for i := 1; i < cnt-1 ; i++ {
    if up && A[i-1] >= A[i] {
        return false
    } 
    if !up && A[i] >= A[i-1] {
        return false
    }
    if A[i] >= A[i+1] {
        //下坡
        up = false
    }
}
return true
}