package {{{name}}}

type TreeNode struct {
  Text string
  Offset int
  Elements []TreeNode
  NodeType string
}  

func (t *TreeNode) String() string {
  return t.Text
}

func (t *TreeNode) GetByNodeType(nodeType string) *TreeNode {
  for _, el := range t.Elements {
    if el.NodeType == nodeType {
      return &el  
    }
  }
  return nil
}
