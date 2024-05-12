// Code generated by MockGen. DO NOT EDIT.
// Source: tui.go
//
// Generated by this command:
//
//	mockgen -source=tui.go -destination=mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	tcell "github.com/gdamore/tcell/v2"
	tui "github.com/kampanosg/lazytest/internal/tui"
	elements "github.com/kampanosg/lazytest/internal/tui/elements"
	state "github.com/kampanosg/lazytest/internal/tui/state"
	models "github.com/kampanosg/lazytest/pkg/models"
	tview "github.com/rivo/tview"
	gomock "go.uber.org/mock/gomock"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// EnableMouse mocks base method.
func (m *MockApplication) EnableMouse(enable bool) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableMouse", enable)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// EnableMouse indicates an expected call of EnableMouse.
func (mr *MockApplicationMockRecorder) EnableMouse(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableMouse", reflect.TypeOf((*MockApplication)(nil).EnableMouse), enable)
}

// QueueUpdateDraw mocks base method.
func (m *MockApplication) QueueUpdateDraw(f func()) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueUpdateDraw", f)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// QueueUpdateDraw indicates an expected call of QueueUpdateDraw.
func (mr *MockApplicationMockRecorder) QueueUpdateDraw(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUpdateDraw", reflect.TypeOf((*MockApplication)(nil).QueueUpdateDraw), f)
}

// SetFocus mocks base method.
func (m *MockApplication) SetFocus(p tview.Primitive) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFocus", p)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetFocus indicates an expected call of SetFocus.
func (mr *MockApplicationMockRecorder) SetFocus(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFocus", reflect.TypeOf((*MockApplication)(nil).SetFocus), p)
}

// SetInputCapture mocks base method.
func (m *MockApplication) SetInputCapture(capture func(*tcell.EventKey) *tcell.EventKey) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInputCapture", capture)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetInputCapture indicates an expected call of SetInputCapture.
func (mr *MockApplicationMockRecorder) SetInputCapture(capture any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInputCapture", reflect.TypeOf((*MockApplication)(nil).SetInputCapture), capture)
}

// SetRoot mocks base method.
func (m *MockApplication) SetRoot(root tview.Primitive, fullscreen bool) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRoot", root, fullscreen)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetRoot indicates an expected call of SetRoot.
func (mr *MockApplicationMockRecorder) SetRoot(root, fullscreen any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoot", reflect.TypeOf((*MockApplication)(nil).SetRoot), root, fullscreen)
}

// Stop mocks base method.
func (m *MockApplication) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockApplicationMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockApplication)(nil).Stop))
}

// MockRunner is a mock of Runner interface.
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner.
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance.
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// RunTest mocks base method.
func (m *MockRunner) RunTest(command string) *models.LazyTestResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTest", command)
	ret0, _ := ret[0].(*models.LazyTestResult)
	return ret0
}

// RunTest indicates an expected call of RunTest.
func (mr *MockRunnerMockRecorder) RunTest(command any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTest", reflect.TypeOf((*MockRunner)(nil).RunTest), command)
}

// MockHandlers is a mock of Handlers interface.
type MockHandlers struct {
	ctrl     *gomock.Controller
	recorder *MockHandlersMockRecorder
}

// MockHandlersMockRecorder is the mock recorder for MockHandlers.
type MockHandlersMockRecorder struct {
	mock *MockHandlers
}

// NewMockHandlers creates a new mock instance.
func NewMockHandlers(ctrl *gomock.Controller) *MockHandlers {
	mock := &MockHandlers{ctrl: ctrl}
	mock.recorder = &MockHandlersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandlers) EXPECT() *MockHandlersMockRecorder {
	return m.recorder
}

// HandleHelpDone mocks base method.
func (m *MockHandlers) HandleHelpDone(a tui.Application, e *elements.Elements) func(int, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleHelpDone", a, e)
	ret0, _ := ret[0].(func(int, string))
	return ret0
}

// HandleHelpDone indicates an expected call of HandleHelpDone.
func (mr *MockHandlersMockRecorder) HandleHelpDone(a, e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleHelpDone", reflect.TypeOf((*MockHandlers)(nil).HandleHelpDone), a, e)
}

// HandleMoveDown mocks base method.
func (m *MockHandlers) HandleMoveDown(e *elements.Elements) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleMoveDown", e)
}

// HandleMoveDown indicates an expected call of HandleMoveDown.
func (mr *MockHandlersMockRecorder) HandleMoveDown(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMoveDown", reflect.TypeOf((*MockHandlers)(nil).HandleMoveDown), e)
}

// HandleMoveUp mocks base method.
func (m *MockHandlers) HandleMoveUp(e *elements.Elements) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleMoveUp", e)
}

// HandleMoveUp indicates an expected call of HandleMoveUp.
func (mr *MockHandlersMockRecorder) HandleMoveUp(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMoveUp", reflect.TypeOf((*MockHandlers)(nil).HandleMoveUp), e)
}

// HandleNodeChanged mocks base method.
func (m *MockHandlers) HandleNodeChanged(e *elements.Elements, s *state.State) func(*tview.TreeNode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNodeChanged", e, s)
	ret0, _ := ret[0].(func(*tview.TreeNode))
	return ret0
}

// HandleNodeChanged indicates an expected call of HandleNodeChanged.
func (mr *MockHandlersMockRecorder) HandleNodeChanged(e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNodeChanged", reflect.TypeOf((*MockHandlers)(nil).HandleNodeChanged), e, s)
}

// HandleResize mocks base method.
func (m *MockHandlers) HandleResize(d tui.ResizeDirection, e *elements.Elements, s *state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleResize", d, e, s)
}

// HandleResize indicates an expected call of HandleResize.
func (mr *MockHandlersMockRecorder) HandleResize(d, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleResize", reflect.TypeOf((*MockHandlers)(nil).HandleResize), d, e, s)
}

// HandleRun mocks base method.
func (m *MockHandlers) HandleRun(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRun", r, a, e, s)
}

// HandleRun indicates an expected call of HandleRun.
func (mr *MockHandlersMockRecorder) HandleRun(r, a, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRun", reflect.TypeOf((*MockHandlers)(nil).HandleRun), r, a, e, s)
}

// HandleRunAll mocks base method.
func (m *MockHandlers) HandleRunAll(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRunAll", r, a, e, s)
}

// HandleRunAll indicates an expected call of HandleRunAll.
func (mr *MockHandlersMockRecorder) HandleRunAll(r, a, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRunAll", reflect.TypeOf((*MockHandlers)(nil).HandleRunAll), r, a, e, s)
}

// HandleRunFailed mocks base method.
func (m *MockHandlers) HandleRunFailed(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRunFailed", r, a, e, s)
}

// HandleRunFailed indicates an expected call of HandleRunFailed.
func (mr *MockHandlersMockRecorder) HandleRunFailed(r, a, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRunFailed", reflect.TypeOf((*MockHandlers)(nil).HandleRunFailed), r, a, e, s)
}

// HandleRunPassed mocks base method.
func (m *MockHandlers) HandleRunPassed(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRunPassed", r, a, e, s)
}

// HandleRunPassed indicates an expected call of HandleRunPassed.
func (mr *MockHandlersMockRecorder) HandleRunPassed(r, a, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRunPassed", reflect.TypeOf((*MockHandlers)(nil).HandleRunPassed), r, a, e, s)
}

// HandleSearchChanged mocks base method.
func (m *MockHandlers) HandleSearchChanged(e *elements.Elements, s *state.State) func(string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleSearchChanged", e, s)
	ret0, _ := ret[0].(func(string))
	return ret0
}

// HandleSearchChanged indicates an expected call of HandleSearchChanged.
func (mr *MockHandlersMockRecorder) HandleSearchChanged(e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSearchChanged", reflect.TypeOf((*MockHandlers)(nil).HandleSearchChanged), e, s)
}

// HandleSearchClear mocks base method.
func (m *MockHandlers) HandleSearchClear(a tui.Application, e *elements.Elements, s *state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleSearchClear", a, e, s)
}

// HandleSearchClear indicates an expected call of HandleSearchClear.
func (mr *MockHandlersMockRecorder) HandleSearchClear(a, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSearchClear", reflect.TypeOf((*MockHandlers)(nil).HandleSearchClear), a, e, s)
}

// HandleSearchDone mocks base method.
func (m *MockHandlers) HandleSearchDone(a tui.Application, e *elements.Elements, s *state.State) func(tcell.Key) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleSearchDone", a, e, s)
	ret0, _ := ret[0].(func(tcell.Key))
	return ret0
}

// HandleSearchDone indicates an expected call of HandleSearchDone.
func (mr *MockHandlersMockRecorder) HandleSearchDone(a, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSearchDone", reflect.TypeOf((*MockHandlers)(nil).HandleSearchDone), a, e, s)
}

// HandleSearchFocus mocks base method.
func (m *MockHandlers) HandleSearchFocus(a tui.Application, e *elements.Elements, s *state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleSearchFocus", a, e, s)
}

// HandleSearchFocus indicates an expected call of HandleSearchFocus.
func (mr *MockHandlersMockRecorder) HandleSearchFocus(a, e, s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSearchFocus", reflect.TypeOf((*MockHandlers)(nil).HandleSearchFocus), a, e, s)
}

// HandleYankNode mocks base method.
func (m *MockHandlers) HandleYankNode(a tui.Application, c tui.Clipboard, e *elements.Elements) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleYankNode", a, c, e)
}

// HandleYankNode indicates an expected call of HandleYankNode.
func (mr *MockHandlersMockRecorder) HandleYankNode(a, c, e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleYankNode", reflect.TypeOf((*MockHandlers)(nil).HandleYankNode), a, c, e)
}

// HandleYankOutput mocks base method.
func (m *MockHandlers) HandleYankOutput(a tui.Application, c tui.Clipboard, e *elements.Elements) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleYankOutput", a, c, e)
}

// HandleYankOutput indicates an expected call of HandleYankOutput.
func (mr *MockHandlersMockRecorder) HandleYankOutput(a, c, e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleYankOutput", reflect.TypeOf((*MockHandlers)(nil).HandleYankOutput), a, c, e)
}

// MockClipboard is a mock of Clipboard interface.
type MockClipboard struct {
	ctrl     *gomock.Controller
	recorder *MockClipboardMockRecorder
}

// MockClipboardMockRecorder is the mock recorder for MockClipboard.
type MockClipboardMockRecorder struct {
	mock *MockClipboard
}

// NewMockClipboard creates a new mock instance.
func NewMockClipboard(ctrl *gomock.Controller) *MockClipboard {
	mock := &MockClipboard{ctrl: ctrl}
	mock.recorder = &MockClipboardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClipboard) EXPECT() *MockClipboardMockRecorder {
	return m.recorder
}

// WriteAll mocks base method.
func (m *MockClipboard) WriteAll(text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteAll", text)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteAll indicates an expected call of WriteAll.
func (mr *MockClipboardMockRecorder) WriteAll(text any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAll", reflect.TypeOf((*MockClipboard)(nil).WriteAll), text)
}
