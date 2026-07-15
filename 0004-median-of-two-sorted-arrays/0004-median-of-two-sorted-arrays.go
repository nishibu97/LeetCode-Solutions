func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	total := m + n
	half := (total + 1) / 2

	const negInf = -1 << 30
	const posInf = 1 << 30

	left, right := 0, m
	for left <= right {
		i := (left + right) / 2
		j := half - i

		aLeft := negInf
		if i > 0 {
			aLeft = nums1[i-1]
		}
		aRight := posInf
		if i < m {
			aRight = nums1[i]
		}

		bLeft := negInf
		if j > 0 {
			bLeft = nums2[j-1]
		}
		bRight := posInf
		if j < n {
			bRight = nums2[j]
		}

		if aLeft > bRight {
			right = i - 1
			continue
		}
		if bLeft > aRight {
			left = i + 1
			continue
		}

		leftMax := aLeft
		if bLeft > leftMax {
			leftMax = bLeft
		}
		if total%2 == 1 {
			return float64(leftMax)
		}

		rightMin := aRight
		if bRight < rightMin {
			rightMin = bRight
		}
		return float64(leftMax+rightMin) / 2.0
	}

	return 0
}