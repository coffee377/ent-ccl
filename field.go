package entcc

import "entgo.io/ent/schema"

const FieldSortAnnotation = "CCSort"

type fieldSort struct {
	Number   int
	Disabled bool
	tail     bool
	desc     bool // 是否降序排序
}

func (s fieldSort) Name() string {
	return FieldSortAnnotation
}

type FieldSortOption func(*fieldSort)

func Sort(num int, options ...FieldSortOption) schema.Annotation {
	f := fieldSort{Number: num}
	for _, apply := range options {
		apply(&f)
	}
	return f
}

func SortTail(num int, options ...FieldSortOption) schema.Annotation {
	return Sort(num, append(options, Tail())...)
}

func Reversed() FieldSortOption {
	return func(f *fieldSort) {
		f.Number = -f.Number
	}
}

func Tail() FieldSortOption {
	return func(f *fieldSort) {
		f.tail = true
	}
}

func Desc(desc bool) FieldSortOption {
	return func(f *fieldSort) {
		f.desc = desc
	}
}

func Disabled() FieldSortOption {
	return func(f *fieldSort) {
		f.Disabled = true
	}
}
