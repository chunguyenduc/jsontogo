package app

import (
	"github.com/chunguyenduc/jsontogo/internal/builder"
	"github.com/chunguyenduc/jsontogo/internal/exporter"
	"github.com/chunguyenduc/jsontogo/internal/importer"
)

type Application struct {
	Importer importer.Importer
	Builder  builder.Builder
	Exporter exporter.Exporter
}

func NewApplication(importer importer.Importer, builder builder.Builder, exporter exporter.Exporter) *Application {
	return &Application{
		Importer: importer,
		Builder:  builder,
		Exporter: exporter,
	}
}

func (a *Application) RunApp() error {
	json, err := a.Importer.Import()
	if err != nil {
		return err
	}

	result, err := a.Builder.Build(json)
	if err != nil {
		return err
	}

	_, err = a.Exporter.Export(result)
	return err
}
