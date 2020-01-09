package utils

func BubbleSort(element []int) {
	if len(element) < 2 {
		return
	}
	for changed := true; changed == true; {
		changed = false
		for i := 0; i < len(element)-1; i++ {
			el0 := element[i]
			el1 := element[i+1]
			if el0 > el1 {
				changed = true
				element[i], element[i+1] = element[i+1], element[i]
			}
		}
	}
}
