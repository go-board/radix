package radix

// Radix is an optimizated trie tree with path compression.
// a string slice {"hello", "hellx", "helxo", "hyxxx", "world"}
// will generate a trie tree like:
//			root
//			/   \
//         w     h
//        /     / \
//       o     e   y
//      /     /     \
//     r     l       x
//    /     / \       \
//   l     l   x       x
//  /     / \   \       \
// d     o   x   o       x
// while radix tree is:
// 			root
//           /  \
//          h    world
//         / \
//        el  y
//        / \  \
//       l   x  xxx
//      / \   \
//     o   x   o
type Radix struct {
	root  *node
	count int64
}

func New() *Radix {
	return &Radix{root: newNode("", "", nil)}
}

func (r *Radix) Insert(key string, value interface{}) (old interface{}, ok bool) {
	return r.root.insertChild(key, key, value)
}

func (r *Radix) Remove(key string) (value interface{}, ok bool) {

}

func (r *Radix) Find(key string) (value interface{}, ok bool) {

}

func (r *Radix) ViewPrefix(prefix string, fn func(string, interface{})) {

}

func (r *Radix) View(fn func(string, interface{})) {
	r.ViewPrefix("", fn)
}
