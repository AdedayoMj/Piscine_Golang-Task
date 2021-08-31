//
package piscine

// Write a function SortWordArr that sorts by ascii (in ascending order) a string slice.

func SortWordArr(array []string) {
	quickSrot2(array, 0, len(array)-1)
}

func quickSrot2(table []string, beg int, end int) {
	if beg < end {
		lockPivo := mudaVariavel2(table, beg, end)
		quickSrot2(table, beg, lockPivo-1)
		quickSrot2(table, lockPivo+1, end)

	}
}

func mudaVariavel2(table []string, beg int, end int) int {
	pivo := table[end]
	i := beg - 1

	for j := beg; j < end; j++ {
		if table[j] <= pivo {
			i++
			aux := table[j]
			table[j] = table[i]
			table[i] = aux
		}
	}

	aux := table[end]
	table[end] = table[i+1]
	table[i+1] = aux

	return i + 1
}

// func main() {
// 	// result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	result := []string{"phSxpfKuDrrR3", "1j8J", "VPFg74wy", "Amsz9hhe", "gH8t", "hDaWXx", "MnE", "3", "9S", "1RCHp", "8se", "KeaN6REk", "uZYc", "j", "xB", "gG76YuIV", "UplWCIs7rpuNv"}
// 	// [phSxpfKuDrrR3 1j8J VPFg74wy Amsz9hhe gH8t hDaWXx MnE  3 9S  1RCHp 8se KeaN6REk uZYc j xB gG76YuIV UplWCIs7rpuNv]
// 	// [1j8J VPFg74wy 9S  1RCHp 8se Amsz9hhe gH8t KeaN6REk uZYc UplWCIs7rpuNv hDaWXx MnE  3 j xB gG76YuIV phSxpfKuDrrR3]

// 	// result := []string{"cDa1e1SSpd9g", "lVxcwqGBOB", "Z", "0Qt3hNL4BUB"}
// 	// [ cDa1e1SSpd9g   lVxcwqGBOB  Z 0Qt3hNL4BUB MUhVCKvlL v M bB12DtXcUVRD5 2op7be9D JkTs y86P0iE62uMg6 fr NDpJ JW7AE]

// 	SortWordArr(result)

// 	fmt.Println(result)
// }
