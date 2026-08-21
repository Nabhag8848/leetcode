func fourSum(nums []int, target int) [][]int {
   hash_set := make([][]int, 0)
   for i:=0; i < len(nums); i++ {
      for j:= i + 1; j < len(nums); j++ {
          hash_map := make(map[int]struct{})
          for k:= j + 1; k < len(nums); k++ {
            needed := target - (nums[i] + nums[j] + nums[k])
                if _, is_exist := hash_map[needed]; is_exist {
                 possible_ans := []int{nums[i], nums[j], nums[k], needed}
                 sort.Ints(possible_ans)
                 is_found := false
                 for _,v := range hash_set {
                    if v[0] == possible_ans[0] && v[1] == possible_ans[1] && v[2] == possible_ans[2] && v[3] == possible_ans[3] {
                        is_found = true
                        break
                    }
                 }

                 if !is_found {
                    hash_set = append(hash_set, possible_ans)
                 }

                 } else {
                    hash_map[nums[k]] = struct{}{}
                }
          }
          
      }
   }

   return hash_set
}

/*  
   x + y + z + w = target
   -(x + y + z) = -w+target

*/