package addtwonumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root1, root2, carry, ret := l1, l2, 0, new(ListNode)
	toggle := true
	for root1 != nil || root2 != nil {
		i1, i2 := 0, 0
		if root1 != nil {
			i1 = root1.Val
			root1 = root1.Next
		}
		if root2 != nil {
			i2 = root2.Val
			root2 = root2.Next
		}
		sum := i1 + i2 + carry
		val := sum % 10
		carry = int(sum / 10)
		if toggle {
			ret.Val = val
			toggle = false
		} else {
			addElement(ret, val)
		}
	}
	if carry > 0 {
		addElement(ret, carry)
	}
	return ret
}

func CreateList(a []int) *ListNode {
	var root *ListNode
	for i, val := range a {
		if i == 0 {
			root = &ListNode{Val: val}
			continue
		}
		addElement(root, val)
	}
	return root
}

func addElement(root *ListNode, i int) *ListNode {
	r := root
	for r.Next != nil {
		r = r.Next
	}
	r.Next = &ListNode{Val: i}
	return r
}

func Display(root *ListNode) {
	if root != nil {
		fmt.Printf("%d -> ", root.Val)
		Display(root.Next)
	}
}
