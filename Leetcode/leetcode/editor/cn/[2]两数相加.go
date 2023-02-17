package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//使用三个指针变量，两个用来遍历原来的两个链表，
//另一个用来从虚拟头节点开始生成新的链表，
//最后返回虚拟头节点的next节点便是我们需要的答案了
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	p1, p2 := l1, l2
	dummy := ListNode{-1, nil}
	p := &dummy
	carry := 0
	for ok := true; ok; ok=(p1!=nil||p2!=nil||carry>0){
		myvar := carry
		if p1 !=nil{
			myvar += p1.Val
			p1 = p1.Next
		}
		if p2!=nil{
			myvar += p2.Val
			p2 = p2.Next
		}
		carry = myvar/10
		p.Next = &ListNode{myvar%10, nil}
		p = p.Next
	}
	return dummy.Next


}

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers2(l1, l2, 0)
}



func addTwoNumbers2(l1 *ListNode, l2 *ListNode, add int) *ListNode {
	if l1 == nil && l2 == nil {
		if add != 0 {
			return &ListNode{1, nil}
		}
		return nil
	} else if l1 != nil && l2 != nil {
		sum := add + l1.Val + l2.Val
		return &ListNode{
			Val:  sum % 10,
			Next: addTwoNumbers2(l1.Next, l2.Next, sum/10),
		}
	} else if l1 != nil && l2 == nil {
		sum := add + l1.Val
		return &ListNode{
			Val: sum%10,
			Next: addTwoNumbers2(l1.Next, nil, sum/10),
		}

	} else if l1 == nil && l2 != nil{
		sum:=add+l2.Val
		return &ListNode{
			Val: sum %10,
			Next:  addTwoNumbers2(nil, l2.Next, sum/10),
		}
	}else{
		return nil
	}
}

//leetcode submit region end(Prohibit modification and deletion)
