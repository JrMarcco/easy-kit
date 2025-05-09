package errs

import (
	"errors"
	"fmt"
)

var (
	ErrSameRBNode   = errors.New("[easy_kit] cannot insert same red-black tree node")
	ErrNodeNotFound = errors.New("[easy_kit] cannot find node in red-black tree")
)

func NilErr(name string) error {
	return fmt.Errorf("[easy-kit] %s is nil", name)
}

func ErrIndexOutOfBounds(length int, index int) error {
	return fmt.Errorf("[easy-kit] index %d out of bounds for length %d", index, length)
}

func ErrInvalidType(want string, got any) error {
	return fmt.Errorf("[easy-kit] invalid type: want %s, got %#v", want, got)
}

func ErrEmptySlice() error {
	return fmt.Errorf("[easy-kit] slice is empty")
}

func ErrInvalidKeyValLen() error {
	return fmt.Errorf("[easy-kit] keys and vals have different lengths")
}
