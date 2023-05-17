package cmd

type Application struct {
	Builder  StructBuilder
	Exporter StructExporter
}

func NewApplication(builder StructBuilder, exporter StructExporter) *Application {
	return &Application{
		Builder:  builder,
		Exporter: exporter,
	}
}

func (a *Application) RunApp() error {
	result, err := a.Builder.Build()
	if err != nil {
		return err
	}

	_, err = a.Exporter.Export(result)
	return err
}
