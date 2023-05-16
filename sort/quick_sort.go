package sort

func QuickSort(input []int64) []int64 {
	result := quickSort(&input, 0, len(input)-1)
	println("%v", result)
	return input
}

func quickSort(input *[]int64, i, j int) []int64 {
	nums := *input
	if len(nums) == 1 || i >= j {
		return nums
	}
	l := i
	r := j
	temp := nums[r]
	for ; l < r; {
		for ; l < r; {
			if nums[l] >= temp && l < r {
				break
			}
			l++
		}
		for ; l < r; {
			if nums[r] < temp && l < r {
				break
			}
			r--
		}
		Swap(input, r, l)
	}
	Swap(input, r, j)
	quickSort(input, i, r-1)
	quickSort(input, r+1, j)
	return nums
}

func Swap(input *[]int64, i, j int) {
	nums := *input
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
