package array

//输入：nums = [1,1,2]
//输出：2, nums = [1,2,_]
//解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
//https://leetcode.cn/problems/remove-duplicates-from-sorted-array

func removeDuplicates(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

//示例 1:
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//https://leetcode.cn/problems/move-zeroes/

func moveZeroes(nums []int) {
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}

//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6] 。
//合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
//https://leetcode.cn/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for index := 0; index < n; index++ {
			nums1[index] = nums2[index]
		}
		return
	}

	if n == 0 {
		return
	}

	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if i < 0 ||  (j >= 0 && nums1[i] <= nums2[j]) {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
}
