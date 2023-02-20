package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//	     Val int
//	     Next *ListNode
//}
//双指针
func reverseList(head *ListNode) *ListNode {
	var p1 *ListNode = nil
	p2 := head
	for (p2 != nil) {
		tmp := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = tmp
	}
	return p1

}




//递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil{
//		return head
//	}
//	last := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return last
//
//}

//leetcode submit region end(Prohibit modification and deletion)
