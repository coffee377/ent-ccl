package entcc

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"go.uber.org/zap"
)

type ccExtension struct {
	entc.DefaultExtension
	logger *zap.Logger
}

func (cc ccExtension) Hooks() []gen.Hook {
	return []gen.Hook{
		FiledSortHook(cc.logger.Sugar()),
	}
}

// Annotations of the extensions.
func (cc ccExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{}
}

// Templates of the extensions.
func (cc ccExtension) Templates() []*gen.Template { return []*gen.Template{} }

// Options of the extensions.
func (cc ccExtension) Options() []entc.Option {
	return []entc.Option{
		entc.FeatureNames(gen.FeatureGlobalID.Name, gen.FeatureModifier.Name),
	}
}

// ExtensionOption is an option for the ccExtension.
type ExtensionOption func(extension *ccExtension)

// NewExtension returns a new Extension configured by opts.
func NewExtension(opts ...ExtensionOption) (entc.Extension, error) {
	e := &ccExtension{}
	for _, opt := range opts {
		opt(e)
	}
	if e.logger == nil {
		e.logger = zap.NewNop()
	}
	return e, nil
}

func WithZapLogger(logger *zap.Logger) ExtensionOption {
	return func(extension *ccExtension) {
		extension.logger = logger
	}
}
