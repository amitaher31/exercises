import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}
	sort.Ints(nums)
	// fmt.Println(nums)
	r := [][]int{}
	for i, _ := range nums {

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		r = twoSum(nums, i, r)

	}

	return r
}

func twoSum(nums []int, i int, r [][]int) [][]int {
	lo := i + 1
	hi := len(nums) - 1

	for lo < hi {
		if hi < len(nums)-1 && nums[hi] == nums[hi+1] {
			hi--
			continue
		}

		if lo > i+1 && nums[lo] == nums[lo-1] {
			lo++
			continue
		}

		if nums[i]+nums[lo]+nums[hi] < 0 {
			lo++
		} else if nums[i]+nums[lo]+nums[hi] > 0 {
			hi--
		} else {
			r1 := []int{nums[i], nums[lo], nums[hi]}
			r = append(r, r1)
			lo++
			hi--

		}
	}
	return r
}