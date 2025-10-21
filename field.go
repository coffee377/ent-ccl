package entcc

import "entgo.io/ent/schema"

const FieldSortAnnotation = "CCSort"

type fieldSort struct {
	Disabled bool
	Number   int
	tail     bool
	desc     bool // 是否降序排序
}

func (s fieldSort) Name() string {
	return FieldSortAnnotation
}

type FieldSortOption func(*fieldSort)

// FieldSort 实体级别控制
func FieldSort(enable bool) schema.Annotation {
	return fieldSort{Disabled: !enable}
}

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

func Disabled(disabled bool) FieldSortOption {
	return func(f *fieldSort) {
		f.Disabled = disabled
	}
}
