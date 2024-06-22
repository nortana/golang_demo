package main
import (
	"fmt"
)

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		//fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	temp.next = newCatNode
	newCatNode.next = head

}


func ListCircleLink(head *CatNode){
	fmt.Println("环形链表的情况如下》》")
	//fmt.Println("dddd>\n",head.no,head.name)
	//return

	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表.....")
		return
	}

	for {
		fmt.Printf("猫的信息为=[id=%d name=%s] ->\n",temp.no,temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}

}

func DelCatNode(head *CatNode, id int) *CatNode{
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}

	if temp.next == head {
		temp.next = nil
		return head
	}

	for {
		//fmt.Println("helper22",helper.next.no)
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//fmt.Println("helper1",helper.no)

	flag := true
	for {
		if temp.next == head {
			break
		}

		if temp.no == id {
			
			if temp == head {
				head = head.next
				//fmt.Println("head",head.no)
			}
			//fmt.Println("helper",helper.no)
			//fmt.Println("temp",temp.no)
			helper.next = temp.next
			return head
		}

		temp = temp.next
		helper = helper.next
	}

	if flag {
		if temp.no == id {
			helper.next = temp.next
		}
	}

	return head

}


func main(){
	head := &CatNode{}

	cat1 := &CatNode{
		no:1,
		name:"tom1",
	}
	cat2 := &CatNode{
		no:2,
		name:"tom2",
	}
	cat3 := &CatNode{
		no:3,
		name:"tom3",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DelCatNode(head,1)

	ListCircleLink(head)
}















