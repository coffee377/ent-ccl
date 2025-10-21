package entcc

import (
	"encoding/json"
	"sort"
	"strings"

	"entgo.io/ent/entc/gen"
	"go.uber.org/zap"
)

func FiledSortHook(logger *zap.SugaredLogger) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				anno, ok := fieldAnno(node.Annotations)
				if !ok {
					continue
				}
				logger.Debugf("正在为表 %s 按字段权重排序", node.Name)
				a := make([]string, len(node.Fields))
				for i, field := range node.Fields {
					a[i] = field.Name
				}
				logger.Debugf("排序前：%s", strings.Join(a, ", "))
				fields := sortFields(node.Fields, anno.Desc)

				// 将排序后的字段切片替换回节点
				node.Fields = fields

				b := make([]string, len(fields))
				for i, field := range fields {
					b[i] = field.Name
				}
				logger.Debugf("排序后：%s", strings.Join(b, ", "))
			}
			return next.Generate(g)
		})
	}
}

func fieldAnno(annotations gen.Annotations) (*FieldSort, bool) {
	if anno, ok := annotations[SortAnnotation]; ok {
		m, ok := anno.(map[string]any)
		if !ok {
			return nil, false
		}
		marshal, err := json.Marshal(m)
		if err != nil {
			return nil, false
		}
		var fieldSort FieldSort
		err = json.Unmarshal(marshal, &fieldSort)
		if err != nil {
			return nil, false
		}
		annotations[SortAnnotation] = &fieldSort
		return &fieldSort, !fieldSort.Disabled
	}
	return nil, false
}

func sortFields(fields []*gen.Field, desc bool) []*gen.Field {
	var (
		headFields []*gen.Field
		midFields  []*gen.Field
		tailFields []*gen.Field
	)
	for _, field := range fields {
		anno, ok := fieldAnno(field.Annotations)
		if ok {
			if anno.Tail {
				tailFields = append(tailFields, field)
			} else {
				headFields = append(headFields, field)
			}
		} else {
			midFields = append(midFields, field)
		}

	}
	sort.Slice(headFields, func(i, j int) bool {
		ai := tailFields[i].Annotations[SortAnnotation].(*FieldSort)
		aj := tailFields[j].Annotations[SortAnnotation].(*FieldSort)
		if desc {
			return ai.Number > aj.Number
		}
		return ai.Number < aj.Number
	})
	// 保持它们在 Schema 中定义的原始顺序
	sort.Slice(midFields, func(i, j int) bool {
		return midFields[i].Position.Index < midFields[j].Position.Index
	})
	sort.Slice(tailFields, func(i, j int) bool {
		ai := tailFields[i].Annotations[SortAnnotation].(*FieldSort)
		aj := tailFields[j].Annotations[SortAnnotation].(*FieldSort)
		if desc {
			return ai.Number > aj.Number
		}
		return ai.Number < aj.Number
	})
	res := append(headFields, midFields...)
	res = append(res, tailFields...)
	return res
}
