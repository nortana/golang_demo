package main
import (
	"fmt"
)


type HeroNode struct {

	no int
	name string
	nikename string
	pre *HeroNode
	next *HeroNode
}


func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	//fmt.Println("head.next=",temp.next == nil)
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}


func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no>newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {

			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("存在相同的no",newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next

		newHeroNode.pre = temp

		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		


		temp.next = newHeroNode




	}

}

func DelHeroNode(head *HeroNode, id int){
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("id 不存在")
	}
}

func ListHeroNode(head *HeroNode){

	temp := head

	if temp.next == nil {
		fmt.Println("空————————————————————")
		return
	}


	for {
		fmt.Printf("[%d, %s, %s] ==>",temp.next.no,
		 temp.next.name,temp.next.nikename)
		 temp =temp.next
		 if temp.next == nil {
			return
		 }
	}
	fmt.Println()
	
}


func ListHeroNode2(head *HeroNode){
	fmt.Println("倒序打印：")

	temp := head

	if temp.next == nil {
		fmt.Println("空————————————————————")
		return
	}

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}


	for {
		fmt.Printf("[%d, %s, %s] ==>",temp.no,
		 temp.name,temp.nikename)
		 temp =temp.pre
		 if temp.pre == nil {
			return
		 }
	}
	fmt.Println()
	
}


func main(){
	hero := &HeroNode{}

	hero1 := &HeroNode{
		no:1,
		name:"宋江",
		nikename:"及时雨",
	}

	hero2 := &HeroNode{
		no:2,
		name:"卢俊义",
		nikename:"玉麒麟",
	}
	hero3 := &HeroNode{
		no:3,
		name:"林冲",
		nikename:"豹子头",
	}
	
	InsertHeroNode(hero,hero1)
	InsertHeroNode(hero,hero2)
	InsertHeroNode(hero,hero3)
	ListHeroNode(hero)

	fmt.Println()
	//DelHeroNode(hero,2)
	//ListHeroNode(hero)
	ListHeroNode2(hero)

}

















