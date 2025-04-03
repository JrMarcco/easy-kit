package tree

import (
	"testing"

	"github.com/JrMarcco/easy_kit"
	"github.com/stretchr/testify/assert"
)

func cmp() easy_kit.Comparator[int] {
	return func(a, b int) int { return a - b }
}

func validRBTree[K any, V any](root *rbNode[K, V]) bool {
	if root.getColor() != black {
		return false
	}

	// count the number of black nodes on the path from the root to the farthest leaf
	cnt := 0
	num := 0
	node := root

	// count the black nodes on the path to the leftmost leaf
	for node != nil {
		if node.getColor() == black {
			cnt++
		}
		node = node.left
	}

	return validRBNode(root, cnt, num)
}

func validRBNode[K any, V any](node *rbNode[K, V], cnt int, num int) bool {
	if node == nil {
		return true
	}

	// red node with red parent is invalid
	if node.getColor() == red && node.parent.getColor() == red {
		return false
	}

	if node.getColor() == black {
		num++
	}

	if node.left == nil && node.right == nil {
		// leaf node
		if num != cnt {
			return false
		}
	}

	return validRBNode(node.left, cnt, num) && validRBNode(node.right, cnt, num)
}

func TestNewRBTree(t *testing.T) {
	tcs := []struct {
		name    string
		cmp     easy_kit.Comparator[int]
		wantRes bool
	}{
		{
			name:    "int cmp",
			cmp:     cmp(),
			wantRes: true,
		}, {
			name:    "nil cmp",
			cmp:     nil,
			wantRes: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rbt := NewRBTree[int, string](tc.cmp)
			assert.Equal(t, tc.wantRes, validRBTree(rbt.root))
		})
	}
}

func TestRBTree_ValidRBTree(t *testing.T) {
	tcs := []struct {
		name    string
		node    *rbNode[int, int]
		wantRes bool
	}{
		{
			name:    "nil",
			node:    nil,
			wantRes: true,
		}, {
			name:    "root with color black",
			node:    &rbNode[int, int]{left: nil, right: nil, color: black},
			wantRes: true,
		}, {
			name:    "root with color red",
			node:    &rbNode[int, int]{left: nil, right: nil, color: red},
			wantRes: false,
		}, {
			name: "root with one child",
			node: &rbNode[int, int]{
				left: &rbNode[int, int]{
					right: nil,
					left:  nil,
					color: red,
				},
				right: nil,
				color: black,
			},
			wantRes: true,
		}, {
			name: "root with two children",
			node: &rbNode[int, int]{
				left: &rbNode[int, int]{
					right: nil,
					left:  nil,
					color: red,
				},
				right: &rbNode[int, int]{
					right: nil,
					left:  nil,
					color: black,
				},
				color: black,
			},
			wantRes: false,
		}, {
			name: "root with grandson (single red node child)",
			node: &rbNode[int, int]{
				left: &rbNode[int, int]{
					right: &rbNode[int, int]{
						right: nil,
						left:  nil,
						color: red,
					},
					left:  nil,
					color: black,
				},
				right: &rbNode[int, int]{
					right: nil,
					left: &rbNode[int, int]{
						right: nil,
						left:  nil,
						color: red,
					},
					color: black,
				},
				color: black,
			},
			wantRes: true,
		}, {
			name: "root with grandson (full red node children)",
			node: &rbNode[int, int]{
				parent: nil,
				key:    7,
				left: &rbNode[int, int]{
					key:   5,
					color: black,
					left: &rbNode[int, int]{
						key:   4,
						color: red,
					},
					right: &rbNode[int, int]{
						key:   6,
						color: red,
					},
				},
				right: &rbNode[int, int]{
					key:   10,
					color: red,
					left: &rbNode[int, int]{
						key:   9,
						color: black,
						left: &rbNode[int, int]{
							key:   8,
							color: red,
						},
					},
					right: &rbNode[int, int]{
						key:   12,
						color: black,
						left: &rbNode[int, int]{
							key:   11,
							color: red,
						},
					},
				},
				color: black,
			},
			wantRes: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantRes, validRBTree(tc.node))
		})
	}
}

func TestRBTree_Insert(t *testing.T) {
	tcs := []struct {
		name        string
		insertNodes []*rbNode[int, int]
		wantRes     bool
		wantErr     error
		wantSize    int64
		wantKeys    []int
		wantVals    []int
	}{
		{
			name:        "insert one node(insert root node)",
			insertNodes: []*rbNode[int, int]{{key: 1, val: 1}},
			wantRes:     true,
			wantErr:     nil,
			wantSize:    1,
			wantKeys:    []int{1},
			wantVals:    []int{1},
		}, {
			name:        "insert two nodes(insert to black parent node)",
			insertNodes: []*rbNode[int, int]{{key: 1, val: 1}, {key: 2, val: 2}},
			wantRes:     true,
			wantErr:     nil,
			wantSize:    2,
			wantKeys:    []int{1, 2},
			wantVals:    []int{1, 2},
		}, {
			name:        "insert multi nodes",
			insertNodes: []*rbNode[int, int]{{key: 1, val: 1}, {key: 2, val: 2}, {key: 3, val: 3}, {key: 4, val: 4}, {key: 5, val: 5}},
			wantRes:     true,
			wantErr:     nil,
			wantSize:    5,
			wantKeys:    []int{1, 2, 3, 4, 5},
			wantVals:    []int{1, 2, 3, 4, 5},
		}, {
			name:        "insert multi desc order nodes",
			insertNodes: []*rbNode[int, int]{{key: 5, val: 5}, {key: 4, val: 4}, {key: 3, val: 3}, {key: 2, val: 2}, {key: 1, val: 1}},
			wantRes:     true,
			wantErr:     nil,
			wantSize:    5,
			wantKeys:    []int{1, 2, 3, 4, 5},
			wantVals:    []int{1, 2, 3, 4, 5},
		}, {
			name:        "insert multi disorder nodes",
			insertNodes: []*rbNode[int, int]{{key: 1, val: 1}, {key: 3, val: 3}, {key: 2, val: 2}, {key: 4, val: 4}, {key: 5, val: 5}},
			wantRes:     true,
			wantErr:     nil,
			wantSize:    5,
			wantKeys:    []int{1, 2, 3, 4, 5},
			wantVals:    []int{1, 2, 3, 4, 5},
		}, {
			name:        "insert same key",
			insertNodes: []*rbNode[int, int]{{key: 1, val: 1}, {key: 1, val: 2}},
			wantErr:     ErrSameRBNode,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rbt := NewRBTree[int, int](cmp())
			for _, node := range tc.insertNodes {
				err := rbt.Insert(node.key, node.val)
				if err != nil {
					assert.Equal(t, tc.wantErr, err)
					return
				}
			}

			assert.Equal(t, tc.wantRes, validRBTree(rbt.root))
			assert.Equal(t, tc.wantSize, int64(rbt.Size()))

			keys, vals := rbt.Kvs()
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantVals, vals)
		})
	}
}
