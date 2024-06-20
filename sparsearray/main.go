package main
import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

//稀疏数组
func main(){
	var chessMap[11][11]int
	chessMap[1][1] = 1 //黑子
	chessMap[2][3] = 2 // 篮子

	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Print(v2,"\t")
		}
		fmt.Println()
	} 

	fmt.Println("----------------------")

	var sparseArr []ValNode

	valNode := ValNode{
		row:11,
		col:11,
		val:0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row:i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
		fmt.Println()
	} 
	fmt.Println("sparseArr = ",sparseArr)

	for i, valNode := range sparseArr {
		fmt.Printf("%d:%d %d %d\n",i,valNode.row,valNode.col,valNode.val)
	}
	


}




