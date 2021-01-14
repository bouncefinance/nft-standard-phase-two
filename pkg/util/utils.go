package util

import (
	"nft_standard/config"
	"strings"
)

func Bool2Int(b bool) int {
	if b{
		return 1
	}else {
		return 0
	}
}

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	tem := arr[i]
	for i < j {
		for i < j && arr[j] >= tem {
			j--
		}
		if i < j && arr[j] < tem {
			arr[i] = arr[j]
			i++
		}

		for i < j && arr[i] < tem {
			i++
		}
		if i < j && arr[i] > tem {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = tem
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}


func StrToLow(str string)(strLow string)  {
	strLow = strings.ToLower(str)
	return
}

func DivideTime(fromBlock, toBlock int64) int64 {
	var section = toBlock - fromBlock + 1
	var time = section / config.EVENT_LOG_SECTION
	if section%config.EVENT_LOG_SECTION > 0 {
		time++
	}
	return time
}