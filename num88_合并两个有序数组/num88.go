package num88

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, n+m)
	i := 0
	j := 0
	k := 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}

	for ; i < m; i++ {
		tmp[k] = nums1[i]
		k++
	}
	for ; j < n; j++ {
		tmp[k] = nums2[j]
		k++
	}

	copy(nums1, tmp)
}
