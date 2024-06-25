package main
import (
	"fmt"
)

func QuickSort(left int, right int, array *[6]int){
	l := left
	r := right
	pivot := array[(left+right)/2]
	temp := 0


	//for 循环的目标是将比pivot 小的数放到 左边
	//比 pivot大的数放到右边
	for ;l<r; {
		for ;array[l] < pivot; {
			l++
		}
		for ;array[r] > pivot;{	
			r--
		 }

		 for l >= r {
			break
		 }
		 temp = array[l]
		 array[l] = array[r]
		 array[r] = temp

		//  if array[l] == pivot {
		// 	r--
		//  }
		//  if array[r] == pivot {
		// 	l++
		//  }

	}
	fmt.Println("left right",left, right)

	if l == r {
		// left++
		// right--
		l++
		r--
 
	}

 
	if left < r {
		QuickSort(left, r, array)
	}
 
	if right > l {
		QuickSort(l, right, array)
	}

}

func main(){
	arr := [6]int{-9, 78, 0, 23, -567, 70}

	fmt.Println(arr)
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)

}