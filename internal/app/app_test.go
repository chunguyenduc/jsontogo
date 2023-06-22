package app

import (
	"errors"
	"fmt"
	"testing"

	"github.com/chunguyenduc/jsontogo/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	app := NewApplication(nil, nil, nil)
	assert.NotNil(t, app)
}

func TestApplication_RunApp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockImporter(ctrl)
	m.
		EXPECT().
		Import().
		Return(nil, errors.New("error"))

	app := NewApplication(m, nil, nil)
	err := app.RunApp()
	assert.NotNil(t, err)

	importer := mocks.NewMockImporter(ctrl)
	importer.
		EXPECT().
		Import().
		Return([]byte("byte"), nil).AnyTimes()

	builder := mocks.NewMockBuilder(ctrl)
	builder.
		EXPECT().
		Build(gomock.Any()).
		Return("", fmt.Errorf("error")).Times(1)

	app = NewApplication(importer, builder, nil)
	err = app.RunApp()
	assert.NotNil(t, err)

	importer.
		EXPECT().
		Import().
		Return([]byte("byte"), nil).AnyTimes()

	builder.
		EXPECT().
		Build([]byte("byte")).
		Return("result", nil).Times(1)

	exporter := mocks.NewMockExporter(ctrl)
	exporter.
		EXPECT().
		Export(gomock.Any()).
		Return(0, fmt.Errorf("error")).AnyTimes()

	app = NewApplication(importer, builder, exporter)
	err = app.RunApp()
	assert.NotNil(t, err)
}
