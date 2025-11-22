package jobs

// Node represents a node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// hasCycle returns true if there is a cycle in the linked list
func hasCycle(head *Node) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // move 1 step
		fast = fast.Next.Next // move 2 steps

		if slow == fast { // they met → cycle exists
			return true
		}
	}

	return false // fast reached the end → no cycle
}
