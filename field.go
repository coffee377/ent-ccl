package entcc

import "entgo.io/ent/schema"

const (
	SortAnnotation = "CCSort"
)

type FieldSort struct {
	Disabled bool // 是否禁用
	Number   int  // 排序数字
	Tail     bool // 是否尾排序
	Desc     bool // 是否降序排序
}

func (s FieldSort) Merge(other schema.Annotation) schema.Annotation {
	var ant FieldSort
	switch other := other.(type) {
	case FieldSort:
		ant = other
	case *FieldSort:
		if other != nil {
			ant = *other
		}
	default:
		return s
	}

	if ant.Disabled {
		s.Disabled = true
	}

	if n := ant.Number; n != 0 {
		s.Number = n
	}

	if ant.Tail {
		s.Tail = true
	}

	if ant.Desc {
		s.Desc = true
	}

	return s
}

func (s FieldSort) Name() string {
	return SortAnnotation
}

// WithFieldSort 实体级别控制,是否启用排序
func WithFieldSort(enable bool) schema.Annotation {
	disabled := !enable
	return FieldSort{Disabled: disabled}
}

// WithFieldDesc 实体级别控制,是否降序排序
func WithFieldDesc(desc bool) schema.Annotation {
	return FieldSort{Desc: desc}
}

func Sort(num int) schema.Annotation {
	return FieldSort{Number: num}
}

func Tail(tailed bool) schema.Annotation {
	return FieldSort{Tail: tailed}
}

func TailSort(num int) schema.Annotation {
	return FieldSort{Number: num, Tail: true}
}

func Disabled(disabled bool) schema.Annotation {
	return FieldSort{Disabled: disabled}
}
