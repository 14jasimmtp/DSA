package main

func quickSort(arr []int) []int{

	if len(arr) < 2{
		return arr
	}

	pivot:=arr[len(arr)-1]

	var left,right []int

	for i:=0;i<len(arr)-1;i++{
		
		if arr[i] < pivot{
			left = append(left,arr[i])
		} else{
			right = append(right,arr[i])
		}
	}

	return append(append(quickSort(left),pivot), quickSort(right)...)
}
