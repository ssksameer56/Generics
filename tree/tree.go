package tree

type Tree[T any] struct {
	Data  T
	Left  *Tree[T]
	Right *Tree[T]
}

func New[T any](item T) (Tree[T], error) {
	node := Tree[T]{
		Data:  item,
		Left:  nil,
		Right: nil,
	}
	return node, nil
}

func (ds *Tree[T]) Inorder() ([]T, error) {
	var ordered []T
	if ds.Left != nil {
		data, err := ds.Left.Inorder()
		if err != nil {
			return ordered, err
		}
		ordered = append(ordered, data...)
	}
	ordered = append(ordered, ds.Data)
	if ds.Right != nil {
		data, err := ds.Right.Inorder()
		if err != nil {
			return ordered, err
		}
		ordered = append(ordered, data...)
	}
	return ordered, nil
}
