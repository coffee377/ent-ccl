package entcc

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"go.uber.org/zap"
)

type cclExtension struct {
	entc.DefaultExtension
	logger *zap.Logger
}

func (ccl cclExtension) Hooks() []gen.Hook {
	return []gen.Hook{
		//FiledSortHook(ccl.logger.Sugar()),
	}
}

// Annotations of the extensions.
func (ccl cclExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{}
}

// Templates of the extensions.
func (ccl cclExtension) Templates() []*gen.Template { return []*gen.Template{} }

// Options of the extensions.
func (ccl cclExtension) Options() []entc.Option { return []entc.Option{} }

// ExtensionOption is an option for the entcc extension.
type ExtensionOption func(extension *cclExtension)

// NewExtension returns a new Extension configured by opts.
func NewExtension(opts ...ExtensionOption) (entc.Extension, error) {
	e := &cclExtension{}
	for _, opt := range opts {
		opt(e)
	}
	return e, nil
}
