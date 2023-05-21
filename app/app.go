package app

type Application struct {
	Importer StructImporter
	Builder  StructBuilder
	Exporter StructExporter
}

func NewApplication(importer StructImporter, builder StructBuilder, exporter StructExporter) *Application {
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
