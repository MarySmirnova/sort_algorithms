package sort

import "math/rand"

//Пузырьковая сортировка по возрастанию
func bubbleSort(ar []int) {
	if len(ar) == 0 {
		return
	}
	count := len(ar)
	test := 0
	for j := 0; j < len(ar)-1; j++ {
		for i := 0; i < count-1; i++ {
			if ar[i] > ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				test++
			}
		}
		if test == 0 {
			break
		}
		test = 0
		count--
	}
}

//Пузырьковая сортировка по убыванию
func bubbleSortReversed(ar []int) {
	if len(ar) == 0 {
		return
	}
	count := len(ar)
	test := 0
	for j := 0; j < len(ar)-1; j++ {
		for i := 0; i < count-1; i++ {
			if ar[i] < ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				test++
			}
		}
		if test == 0 {
			break
		}
		test = 0
		count--
	}
}

//Рекурсивная пузырьковая сортировка по возрастанию
func bubbleSortRecursive(ar []int) {
	count := 0
	if len(ar) < 2 {
		return
	}
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i], ar[i+1] = ar[i+1], ar[i]
			count++
		}
	}
	j := len(ar) - 1
	if count > 0 {
		bubbleSortRecursive(ar[:j])
	}
}

//Сортировка выбором,
//поиск максимальных элементов и перемещение их в конец
func selectionSort(ar []int) {
	lenAr := len(ar)
	var maxVal, maxIndex int

	for i := 0; i < len(ar); i++ {
		maxVal, maxIndex = ar[lenAr-1], lenAr-1

		for j := 0; j < lenAr-1; j++ {
			if ar[j] > maxVal {
				maxVal, maxIndex = ar[j], j
			}
		}
		if maxIndex != lenAr-1 {
			ar[maxIndex], ar[lenAr-1] = ar[lenAr-1], ar[maxIndex]
		}
		lenAr--
	}
}

//Сортировка выбором,
//поиск минимальных элементов и перемещение их в начало
func selectionSortByMin(ar []int) {
	count := 0
	var minVal, minIndex int

	for i := count; i < len(ar); i++ {
		minVal, minIndex = ar[count], count

		for j := len(ar) - 1; j > count; j-- {
			if ar[j] < minVal {
				minVal, minIndex = ar[j], j
			}
		}
		if minIndex != 0 {
			ar[minIndex], ar[count] = ar[count], ar[minIndex]
		}
		count++
	}
}

//Двусторонняя сортировка выбором
func SelectionSortDuble(ar []int) {
	count := 0
	lenAr := len(ar)
	var minVal, minIndex, maxVal, maxIndex int

	for i := count; i < lenAr; i++ {
		minVal, minIndex = ar[count], count
		maxVal, maxIndex = ar[lenAr-1], lenAr-1

		for j := count + 1; j < lenAr-1; j++ {
			if ar[j] < minVal {
				minVal, minIndex = ar[j], j
			}
			if ar[j] > maxVal {
				maxVal, maxIndex = ar[j], j
			}
		}
		if minIndex != count {
			if maxIndex == count {
				maxIndex = minIndex
			}
			ar[minIndex], ar[count] = ar[count], ar[minIndex]
		}
		if maxIndex != lenAr-1 {
			ar[maxIndex], ar[lenAr-1] = ar[lenAr-1], ar[maxIndex]
		}
		lenAr--
		count++
	}
}

//Сортировка вставкой
func insertionSort(ar []int) {
	lenAr := len(ar)
	count := 0
	var val int

	for i := count + 1; i < lenAr; i++ {
		val = ar[count+1]
		for j := count; j >= 0; j-- {
			if val < ar[j] {
				ar[j+1] = ar[j]
				if j == 0 {
					ar[j] = val
				}
				continue
			}
			ar[j+1] = val
			break
		}
		count++
	}
}

//Сортировка слиянием
func MergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	centre := len(ar) / 2
	firstAr := ar[:centre]
	secondAr := ar[centre:]
	totalAr := []int{}
	var el int

	firstAr = MergeSort(firstAr)
	secondAr = MergeSort(secondAr)

	for {
		if ((len(firstAr) > 0 && len(secondAr) > 0) && (firstAr[0] < secondAr[0])) || (len(secondAr) == 0 && len(firstAr) > 0) {
			el, firstAr = firstAr[0], firstAr[1:]
			totalAr = append(totalAr, el)
		} else if ((len(firstAr) > 0 && len(secondAr) > 0) && (firstAr[0] > secondAr[0])) || (len(firstAr) == 0 && len(secondAr) > 0) {
			el, secondAr = secondAr[0], secondAr[1:]
			totalAr = append(totalAr, el)
		} else if (len(firstAr) > 0 && len(secondAr) > 0) && (firstAr[0] == secondAr[0]) {
			el, firstAr = firstAr[0], firstAr[1:]
			totalAr = append(totalAr, el)
			el, secondAr = secondAr[0], secondAr[1:]
			totalAr = append(totalAr, el)
		} else {
			break
		}
	}
	return totalAr
}

//Быстрая сортировка
func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	sup := rand.Intn(len(ar))
	count, el := 0, 0

	ar[len(ar)-1], ar[sup] = ar[sup], ar[len(ar)-1]
	for i := range ar {
		if ar[i] < ar[len(ar)-1] {
			ar[i], ar[el] = ar[el], ar[i]
			el++
		}
		if ar[i] == ar[len(ar)-1] {
			count++
		}
	}
	ar[len(ar)-1], ar[el] = ar[el], ar[len(ar)-1]

	if count == len(ar) {
		return
	}

	quickSort(ar[:el])
	quickSort(ar[el:])
}
