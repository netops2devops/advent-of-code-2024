package main

func SimilarityScore() int {
	total := 0
	leftList, rightList := cleanData()
	for _, left := range leftList {
		count := 0 // number of appearances in right side list
		for _, right := range rightList {
			if left == right {
				count++
			}
		}
		total += left * count
	}
	return total
}
