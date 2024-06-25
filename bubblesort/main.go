package main
import (
	"fmt"
)

func main(){
	var arr [5]int = [5]int{19,4,20,50,17}
	for j := 0; j<len(arr)-1; j++{
		for i := 0; i < len(arr)-j-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
