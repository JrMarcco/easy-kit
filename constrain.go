package easykit

// RealNumber is a constraint that matches any integer or floating-point type.
type RealNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Number is a constraint that matches any numeric type.
type Number interface {
	RealNumber | ~complex64 | ~complex128
}
