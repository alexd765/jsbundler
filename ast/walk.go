package ast

//WalkTo nodes with an interesting type.
func (n *Node) WalkTo(types map[string]struct{}) []*Node {
	if _, ok := types[n.Type]; ok {
		return []*Node{n}
	}

	var hits []*Node
	for _, child := range n.Children {
		nodes := child.WalkTo(types)
		if nodes != nil {
			hits = append(hits, nodes...)
		}
	}
	return hits
}
