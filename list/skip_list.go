package list

import (
	easykit "github.com/JrMarcco/easy-kit"
	"github.com/JrMarcco/easy-kit/internal/list"
)

type SkipList[T any] struct {
	skipList *list.SkipList[T]
}

func (sl *SkipList[T]) Insert(val T) {
	sl.skipList.Insert(val)
}

func (sl *SkipList[T]) Delete(target T) bool {
	return sl.skipList.Delete(target)
}

func (sl *SkipList[T]) Exists(target T) bool {
	return sl.skipList.Exists(target)
}

func (sl *SkipList[T]) Get(index int) (T, bool) {
	return sl.skipList.GetByIndex(index)
}

func (sl *SkipList[T]) Peek() (T, bool) {
	return sl.skipList.Peek()
}

func (sl *SkipList[T]) Len() int {
	return sl.skipList.Len()
}

func (sl *SkipList[T]) ToSlice() []T {
	return sl.skipList.ToSlice()
}

func NewSkipList[T any](cmp easykit.Comparator[T]) *SkipList[T] {
	return &SkipList[T]{
		skipList: list.NewSkipList[T](cmp),
	}
}
