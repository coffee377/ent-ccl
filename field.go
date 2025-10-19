package entcc

import "entgo.io/ent/schema"

const FieldSortAnnotation = "CCSortField"

type fieldSort struct {
	Number   int
	Disabled bool
}

func (s fieldSort) Name() string {
	return FieldSortAnnotation
}

type FieldSortOption func(*fieldSort)

func Sort(num uint, options ...FieldSortOption) schema.Annotation {
	f := fieldSort{Number: int(num)}
	for _, apply := range options {
		apply(&f)
	}
	return f
}

func SortReverse(num uint, options ...FieldSortOption) schema.Annotation {
	return Sort(num, append(options, Reversed())...)
}

func Reversed() FieldSortOption {
	return func(f *fieldSort) {
		f.Number = -f.Number
	}
}

func Disabled() FieldSortOption {
	return func(f *fieldSort) {
		f.Disabled = true
	}
}
