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
				logger.Debugf("正在为表 %s 按自动权重排序", node.Name)
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

type annotationInfo struct {
	Disabled bool `json:"Disabled"` // 是否禁用排序
	Number   int  `json:"Number"`   // 排序权重
	Tail     bool `json:"Tail"`     // 是否放在最后
	Desc     bool `json:"desc"`     // 是否降序
}

func fieldAnno(annotations gen.Annotations) (*annotationInfo, bool) {
	if anno, ok := annotations[SortAnnotation]; ok {
		m, ok := anno.(map[string]any)
		if !ok {
			return nil, false
		}
		marshal, err := json.Marshal(m)
		if err != nil {
			return nil, false
		}
		var o annotationInfo
		err = json.Unmarshal(marshal, &o)
		if err != nil {
			return nil, false
		}
		annotations[SortAnnotation] = &o
		return &o, !o.Disabled
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
		if desc {
			return headFields[i].Name > headFields[j].Name
		}
		return headFields[i].Name < headFields[j].Name
	})
	// 保持它们在 Schema 中定义的原始顺序
	sort.Slice(midFields, func(i, j int) bool {
		return midFields[i].Position.Index < midFields[j].Position.Index
	})
	sort.Slice(tailFields, func(i, j int) bool {
		ai := tailFields[i].Annotations[SortAnnotation].(*annotationInfo)
		aj := tailFields[j].Annotations[SortAnnotation].(*annotationInfo)
		if desc {
			return ai.Number > aj.Number
		}
		return ai.Number < aj.Number
	})
	result := append(headFields, midFields...)
	result = append(result, tailFields...)
	return result
}
