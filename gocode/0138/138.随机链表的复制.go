package l0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	newHead := &Node{}
	curNew := newHead
	cur2CurNew := make(map[*Node]*Node)
	cur2CurNew[nil] = nil

	for cur != nil {
		curNew.Next = &Node{
			Val:    cur.Val,
			Next:   cur.Next,
			Random: cur.Random,
		}
		cur2CurNew[cur] = curNew.Next
		cur = cur.Next
		curNew = curNew.Next
	}

	curNew = newHead.Next
	for curNew != nil {
		curNew.Random = cur2CurNew[curNew.Random]
		curNew = curNew.Next
	}
	return newHead.Next
}
