package app

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	app := NewApplication(nil, nil, nil)
	assert.NotNil(t, app)
}

func TestApplication_RunApp(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockStructImporter(ctrl)
	m.
		EXPECT().
		Import().
		Return(nil, errors.New("error"))

	app := NewApplication(m, nil, nil)
	err := app.RunApp()
	assert.NotNil(t, err)

	importer := NewMockStructImporter(ctrl)
	importer.
		EXPECT().
		Import().
		Return([]byte("byte"), nil).AnyTimes()

	builder := NewMockBuilder(ctrl)
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

	exporter := NewMockStructExporter(ctrl)
	exporter.
		EXPECT().
		Export(gomock.Any()).
		Return(0, fmt.Errorf("error")).AnyTimes()

	app = NewApplication(importer, builder, exporter)
	err = app.RunApp()
	assert.NotNil(t, err)
}

// MockStructImporter is a mock of StructImporter interface.
type MockStructImporter struct {
	ctrl     *gomock.Controller
	recorder *MockStructImporterMockRecorder
}

// MockStructImporterMockRecorder is the mock recorder for MockStructImporter.
type MockStructImporterMockRecorder struct {
	mock *MockStructImporter
}

// NewMockStructImporter creates a new mock instance.
func NewMockStructImporter(ctrl *gomock.Controller) *MockStructImporter {
	mock := &MockStructImporter{ctrl: ctrl}
	mock.recorder = &MockStructImporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStructImporter) EXPECT() *MockStructImporterMockRecorder {
	return m.recorder
}

// Import mocks base method.
func (m *MockStructImporter) Import() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Import indicates an expected call of Import.
func (mr *MockStructImporterMockRecorder) Import() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockStructImporter)(nil).Import))
}

// MockStructExporter is a mock of StructExporter interface.
type MockStructExporter struct {
	ctrl     *gomock.Controller
	recorder *MockStructExporterMockRecorder
}

// MockStructExporterMockRecorder is the mock recorder for MockStructExporter.
type MockStructExporterMockRecorder struct {
	mock *MockStructExporter
}

// NewMockStructExporter creates a new mock instance.
func NewMockStructExporter(ctrl *gomock.Controller) *MockStructExporter {
	mock := &MockStructExporter{ctrl: ctrl}
	mock.recorder = &MockStructExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStructExporter) EXPECT() *MockStructExporterMockRecorder {
	return m.recorder
}

// Export mocks base method.
func (m *MockStructExporter) Export(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Export indicates an expected call of Export.
func (mr *MockStructExporterMockRecorder) Export(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockStructExporter)(nil).Export), arg0)
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockBuilder) Build(arg0 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockBuilderMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build), arg0)
}
