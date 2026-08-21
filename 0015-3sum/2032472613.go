func threeSum(nums []int) [][]int {
   hash_set := make([][]int, 0)
   for i:=0; i < len(nums); i++ {
      hash_map := make(map[int]struct{})
      for j:= i + 1; j< len(nums); j++ {
          target := -(nums[i] + nums[j])
          if _, is_exist := hash_map[target]; is_exist {
             possible_ans := []int{nums[i], nums[j], target}
             sort.Ints(possible_ans)
             is_found := false
             for _,v := range hash_set {
                if v[0] == possible_ans[0] && v[1] == possible_ans[1] && v[2] == possible_ans[2] {
                    is_found = true
                    break
                }
             }

             if !is_found {
                hash_set = append(hash_set, possible_ans)
             }
             
          } else {
             hash_map[nums[j]] = struct{}{}
          }
      }
   }

   return hash_set
}

/*
x + y = k
x - k = -y


x + y + z = 0
x + y = -z

    -1 2
     0 1
     1 1
     2 1
    -4 1

*/
