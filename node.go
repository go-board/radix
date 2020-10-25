package radix

type node struct {
	key      string
	fullKey  string
	value    interface{}
	children []*node
}

func newNode(key string, fullKey string, value interface{}) *node {
	return &node{key: key, fullKey: fullKey, value: value}
}

func (n *node) insertChild(key string, fullKey string, value interface{}) (old interface{}, ok bool) {

}

func (n *node) removeChild(key string, fullKey string) (value interface{}, ok bool) {

}

func (n *node) hasValue() bool {
	return n.value != nil
}
