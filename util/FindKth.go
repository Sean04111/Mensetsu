package util

//找出两个有序数组的第 k 个元素

func GetKth(nums1, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	index1, index2 := 0, 0
	for {
		if index1 == n1 {
			return nums2[index2+k-1]
		}
		if index2 == n2 {
			return nums1[k+index2-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		l1 := min(index1+half, n1) - 1
		l2 := min(index2+half, n2) - 1
		if nums1[l1] <= nums2[l2] {
			k -= l1 - index1 + 1
			index1 = l1 + 1
		} else {
			k -= l2 - index2 + 1
			index2 = l2 + 1
		}
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//1 2 3 4
//5 6 7 8 9 10
