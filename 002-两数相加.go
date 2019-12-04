package main

import "git.forms.io/universe/dts/common/log"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l2 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	l3 := addTwoNumbers(l1, l2)
	log.Infof("%++v", l3)
}

// 将长度较短的链表在末尾补零使得两个连表长度相等，再一个一个元素对其相加（考虑进位）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var len1 = 1
	var len2 = 1
	p := l1
	q := l2
	for p.Next != nil {
		p = p.Next
		len1++
	}
	for q.Next != nil {
		q = q.Next
		len2++
	}
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			n := &ListNode{
				Val: 0,
			}
			q.Next = n
			q = q.Next
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			n := &ListNode{
				Val: 0,
			}
			p.Next = n
			p = p.Next
		}
	}
	p = l1
	q = l2
	r := &ListNode{}
	s := r
	var i int
	var j int
	for p != nil && q != nil {
		i = p.Val + q.Val + j
		s.Next = &ListNode{
			Val:  i % 10,
			Next: nil,
		}
		p = p.Next
		q = q.Next
		s = s.Next
		if i > 9 {
			j = 1
		} else {
			j = 0
		}
	}
	if j == 1 {
		s.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return r.Next
}

// 无需补齐，直接记录每次相加的结果
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	s := r
	var sum int
	var carry bool
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry {
			sum += 1
		}
		s.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		s = s.Next
		if sum > 9 {
			carry = true
		} else {
			carry = false
		}
	}
	if carry {
		s.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return r.Next
}
