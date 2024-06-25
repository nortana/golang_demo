package main
import (
	"fmt"
)

func SetWay(myMap *[8][7]int, i int, j int) bool {

	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 {

			myMap[i][j] = 2

			if SetWay(myMap, i-1,j) { 
				return true
			} else if SetWay(myMap, i+1,j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else {
				myMap[i][j] = 3
				return false
			}

		} else {
			return false
		}
	}


}

func main(){
	//先创建一个二位数组，模拟迷宫
	//规则
	//1、如果元素的值为1，就是墙
	//2、如果元素的值为0，是没有走过的点
	//3、如果元素的值为2，就是一格通路
	//4、如果元素的值为3，是走过的点，但是走不通
	//
	//

	var myMap [8][7]int

	for i:= 0; i< 7;i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1

	}
	for i:= 0; i< 7;i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1

	}

	myMap[3][1] = 1
	myMap[3][2] = 1

	SetWay(&myMap,1,1)

	for i := 0;i<8;i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j]," ")
		}
		fmt.Println()
	}

	

}





















