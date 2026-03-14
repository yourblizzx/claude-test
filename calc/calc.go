package calc

// ComputeValues возвращает слайс результатов вычисления 100/i для i от 1 до n.
// Возвращает ошибку, если n <= 0.
func ComputeValues(n int) ([]int, error) {
	if n <= 0 {
		return nil, ErrNonPositive
	}

	results := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		results = append(results, 100/i)
	}
	return results, nil
}

// Sum возвращает сумму элементов слайса.
func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
