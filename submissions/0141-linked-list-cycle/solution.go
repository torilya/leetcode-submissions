/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    for fastPointer, slowPointer := head, head; fastPointer != nil && fastPointer.Next != nil; {
        fastPointer = fastPointer.Next.Next
        slowPointer = slowPointer.Next

        if fastPointer == slowPointer {
            return true
        }
    }

    return false
}

