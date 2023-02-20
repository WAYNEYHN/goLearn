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
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = startReserse(l1).Next
	l2 = startReserse(l2).Next
}

func startReserse(mylist *ListNode) *ListNode{
	virtualNode := ListNode{-1, nil}
	p1 := virtualNode
	return reverse(&p1, mylist)
}

func reverse(p *ListNode, p1 *ListNode) *ListNode{
	if p1.Next == nil {
		return nil
	} else{
		p.Next = p1
		return reverse(p.Next, p1.Next)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
