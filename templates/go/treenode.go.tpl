package {{name}}

type TreeNode interface {
	GetText() string
}

type BaseNode struct {
	Text string
}

func (n *BaseNode) GetText() string {
	return n.Text
}
