package entcc

import "entgo.io/ent/schema"

const (
	SortAnnotation   = "CCSort"
	SortDisabledName = "Disabled"
	SortNumberName   = "Number"
	SortTailName     = "Tail"
	SortDescName     = "Desc"
)

type fieldSort struct {
	Disabled bool
	Number   int
	Tail     bool
	Desc     bool // 是否降序排序
}

func (s fieldSort) Name() string {
	return SortAnnotation
}

type FieldSortOption func(*fieldSort)

// WithFieldSort 实体级别控制
func WithFieldSort(enable bool) schema.Annotation {
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
		f.Tail = true
	}
}

func Desc(desc bool) FieldSortOption {
	return func(f *fieldSort) {
		f.Desc = desc
	}
}

func Disabled(disabled bool) FieldSortOption {
	return func(f *fieldSort) {
		f.Disabled = disabled
	}
}
