package main
import (
	"fmt"
	"errors"
	"strconv"
)

type Stack struct{
	MaxTop int
	Top int
	arr [20]int
}

func (this *Stack) Push(val int) (err error){

	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	this.Top++
	this.arr[this.Top] =val
	return
}

func (this *Stack) Pop() (val int, err error) {

	if this.Top == - 1 {
		fmt.Println("stack empty")
		return 0,errors.New("stack empty")
	}

	val = this.arr[this.Top]

	this.Top--
	return val,nil


}

func (this *Stack) List(){
	if this.Top == -1 {
		fmt.Println("stack emtyp")
		return
	}

	//curTop := this.Top
	fmt.Println("栈的情况如下：：")
	for i := this.Top;i>=0;i-- {
		fmt.Printf("arr[%d]=%d\n",i,this.arr[i])
	}
}

func (this *Stack) IsOper(val int) bool{
	if val == 42 || val == 43 || val == 45 || val ==47 {
		return true
	} else {
		return false
	}
	return false
}

func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
		case 42 :
			res = num1 * num2
		case 43 :
			res = num2 + num1
		case 45:
			res = num2 - num1
		case 47 :
			res = num2 / num1
		default:
			fmt.Println("运算符错误")
	}
	return res
}


func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	 return res

}


func main(){

	numStack := &Stack{
		MaxTop:20,
		Top:-1,
	}

	operStack := &Stack{
		MaxTop:20,
		Top:-1,
	}

	exp := "30+30*6-4"

	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		ch := exp[index:index+1]
		//ASCiI码
		temp := int([]byte(ch)[0])

		if operStack.IsOper(temp) {

			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) > 
					operStack.Priority(temp) {
						num1, _ = numStack.Pop()
						num2, _ = numStack.Pop()
						oper, _ = operStack.Pop()
						result = operStack.Cal(num1,num2,oper)
						numStack.Push(result)
						operStack.Push(temp)

					}  else {
						operStack.Push(temp)
					}
			}

		} else {

			keepNum +=ch
			if index == len(exp) - 1 {
				val, _ := strconv.ParseInt(keepNum,10,64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int([]byte(exp[index+1:index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum,10,64)
					numStack.Push(int(val))
					keepNum = ""

				}
			}

			// val, _ := strconv.ParseInt(ch,10,64)
			// numStack.Push(int(val))
		}

		if index + 1 == len(exp) {
			break
		}
		index++
	}

	for {
		if operStack.Top == -1 {
			break
		}

		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1,num2,oper)
		numStack.Push(result)

	}

	res, _ := numStack.Pop()
	fmt.Println("表达式结果》》》 ",exp," = ",res)
	

}





