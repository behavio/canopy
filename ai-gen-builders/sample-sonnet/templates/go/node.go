
type {{name}} struct {
	Text     string
	Offset   int
	Elements []interface{}
}

func (n *{{name}}) SetText(text string)      { n.Text = text }
func (n *{{name}}) SetOffset(offset int)     { n.Offset = offset }
func (n *{{name}}) SetElements(elements []interface{}) { n.Elements = elements }

func makeNode(text string, offset int, elements []interface{}, action func, nodeType func() *Node) *Node {
	node := nodeType()
	node.SetText(text)
	node.SetOffset(offset)
	node.SetElements(elements)

	if action != nil {
		result := action(text, offset, elements)
		if result != nil {
			node = result.(*Node)
		}
	}

	return node
}
