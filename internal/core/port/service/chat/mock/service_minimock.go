// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mock

//go:generate minimock -i github.com/pillarion/practice-chat-server/internal/core/port/service/chat.Service -o service_minimock.go -n ServiceMock -p mock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	modelChat "github.com/pillarion/practice-chat-server/internal/core/model/chat"
	modelMessage "github.com/pillarion/practice-chat-server/internal/core/model/message"
)

// ServiceMock implements chat.Service
type ServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateChat          func(ctx context.Context, username []modelChat.Username) (i1 int64, err error)
	inspectFuncCreateChat   func(ctx context.Context, username []modelChat.Username)
	afterCreateChatCounter  uint64
	beforeCreateChatCounter uint64
	CreateChatMock          mServiceMockCreateChat

	funcDeleteChat          func(ctx context.Context, id int64) (err error)
	inspectFuncDeleteChat   func(ctx context.Context, id int64)
	afterDeleteChatCounter  uint64
	beforeDeleteChatCounter uint64
	DeleteChatMock          mServiceMockDeleteChat

	funcSendMessage          func(ctx context.Context, message *modelMessage.Message) (err error)
	inspectFuncSendMessage   func(ctx context.Context, message *modelMessage.Message)
	afterSendMessageCounter  uint64
	beforeSendMessageCounter uint64
	SendMessageMock          mServiceMockSendMessage
}

// NewServiceMock returns a mock for chat.Service
func NewServiceMock(t minimock.Tester) *ServiceMock {
	m := &ServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateChatMock = mServiceMockCreateChat{mock: m}
	m.CreateChatMock.callArgs = []*ServiceMockCreateChatParams{}

	m.DeleteChatMock = mServiceMockDeleteChat{mock: m}
	m.DeleteChatMock.callArgs = []*ServiceMockDeleteChatParams{}

	m.SendMessageMock = mServiceMockSendMessage{mock: m}
	m.SendMessageMock.callArgs = []*ServiceMockSendMessageParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mServiceMockCreateChat struct {
	mock               *ServiceMock
	defaultExpectation *ServiceMockCreateChatExpectation
	expectations       []*ServiceMockCreateChatExpectation

	callArgs []*ServiceMockCreateChatParams
	mutex    sync.RWMutex
}

// ServiceMockCreateChatExpectation specifies expectation struct of the Service.CreateChat
type ServiceMockCreateChatExpectation struct {
	mock    *ServiceMock
	params  *ServiceMockCreateChatParams
	results *ServiceMockCreateChatResults
	Counter uint64
}

// ServiceMockCreateChatParams contains parameters of the Service.CreateChat
type ServiceMockCreateChatParams struct {
	ctx      context.Context
	username []modelChat.Username
}

// ServiceMockCreateChatResults contains results of the Service.CreateChat
type ServiceMockCreateChatResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for Service.CreateChat
func (mmCreateChat *mServiceMockCreateChat) Expect(ctx context.Context, username []modelChat.Username) *mServiceMockCreateChat {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ServiceMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ServiceMockCreateChatExpectation{}
	}

	mmCreateChat.defaultExpectation.params = &ServiceMockCreateChatParams{ctx, username}
	for _, e := range mmCreateChat.expectations {
		if minimock.Equal(e.params, mmCreateChat.defaultExpectation.params) {
			mmCreateChat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateChat.defaultExpectation.params)
		}
	}

	return mmCreateChat
}

// Inspect accepts an inspector function that has same arguments as the Service.CreateChat
func (mmCreateChat *mServiceMockCreateChat) Inspect(f func(ctx context.Context, username []modelChat.Username)) *mServiceMockCreateChat {
	if mmCreateChat.mock.inspectFuncCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("Inspect function is already set for ServiceMock.CreateChat")
	}

	mmCreateChat.mock.inspectFuncCreateChat = f

	return mmCreateChat
}

// Return sets up results that will be returned by Service.CreateChat
func (mmCreateChat *mServiceMockCreateChat) Return(i1 int64, err error) *ServiceMock {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ServiceMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ServiceMockCreateChatExpectation{mock: mmCreateChat.mock}
	}
	mmCreateChat.defaultExpectation.results = &ServiceMockCreateChatResults{i1, err}
	return mmCreateChat.mock
}

// Set uses given function f to mock the Service.CreateChat method
func (mmCreateChat *mServiceMockCreateChat) Set(f func(ctx context.Context, username []modelChat.Username) (i1 int64, err error)) *ServiceMock {
	if mmCreateChat.defaultExpectation != nil {
		mmCreateChat.mock.t.Fatalf("Default expectation is already set for the Service.CreateChat method")
	}

	if len(mmCreateChat.expectations) > 0 {
		mmCreateChat.mock.t.Fatalf("Some expectations are already set for the Service.CreateChat method")
	}

	mmCreateChat.mock.funcCreateChat = f
	return mmCreateChat.mock
}

// When sets expectation for the Service.CreateChat which will trigger the result defined by the following
// Then helper
func (mmCreateChat *mServiceMockCreateChat) When(ctx context.Context, username []modelChat.Username) *ServiceMockCreateChatExpectation {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ServiceMock.CreateChat mock is already set by Set")
	}

	expectation := &ServiceMockCreateChatExpectation{
		mock:   mmCreateChat.mock,
		params: &ServiceMockCreateChatParams{ctx, username},
	}
	mmCreateChat.expectations = append(mmCreateChat.expectations, expectation)
	return expectation
}

// Then sets up Service.CreateChat return parameters for the expectation previously defined by the When method
func (e *ServiceMockCreateChatExpectation) Then(i1 int64, err error) *ServiceMock {
	e.results = &ServiceMockCreateChatResults{i1, err}
	return e.mock
}

// CreateChat implements chat.Service
func (mmCreateChat *ServiceMock) CreateChat(ctx context.Context, username []modelChat.Username) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreateChat.beforeCreateChatCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateChat.afterCreateChatCounter, 1)

	if mmCreateChat.inspectFuncCreateChat != nil {
		mmCreateChat.inspectFuncCreateChat(ctx, username)
	}

	mm_params := ServiceMockCreateChatParams{ctx, username}

	// Record call args
	mmCreateChat.CreateChatMock.mutex.Lock()
	mmCreateChat.CreateChatMock.callArgs = append(mmCreateChat.CreateChatMock.callArgs, &mm_params)
	mmCreateChat.CreateChatMock.mutex.Unlock()

	for _, e := range mmCreateChat.CreateChatMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreateChat.CreateChatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateChat.CreateChatMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateChat.CreateChatMock.defaultExpectation.params
		mm_got := ServiceMockCreateChatParams{ctx, username}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateChat.t.Errorf("ServiceMock.CreateChat got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateChat.CreateChatMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateChat.t.Fatal("No results are set for the ServiceMock.CreateChat")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreateChat.funcCreateChat != nil {
		return mmCreateChat.funcCreateChat(ctx, username)
	}
	mmCreateChat.t.Fatalf("Unexpected call to ServiceMock.CreateChat. %v %v", ctx, username)
	return
}

// CreateChatAfterCounter returns a count of finished ServiceMock.CreateChat invocations
func (mmCreateChat *ServiceMock) CreateChatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateChat.afterCreateChatCounter)
}

// CreateChatBeforeCounter returns a count of ServiceMock.CreateChat invocations
func (mmCreateChat *ServiceMock) CreateChatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateChat.beforeCreateChatCounter)
}

// Calls returns a list of arguments used in each call to ServiceMock.CreateChat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateChat *mServiceMockCreateChat) Calls() []*ServiceMockCreateChatParams {
	mmCreateChat.mutex.RLock()

	argCopy := make([]*ServiceMockCreateChatParams, len(mmCreateChat.callArgs))
	copy(argCopy, mmCreateChat.callArgs)

	mmCreateChat.mutex.RUnlock()

	return argCopy
}

// MinimockCreateChatDone returns true if the count of the CreateChat invocations corresponds
// the number of defined expectations
func (m *ServiceMock) MinimockCreateChatDone() bool {
	for _, e := range m.CreateChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateChat != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateChatInspect logs each unmet expectation
func (m *ServiceMock) MinimockCreateChatInspect() {
	for _, e := range m.CreateChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ServiceMock.CreateChat with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		if m.CreateChatMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ServiceMock.CreateChat")
		} else {
			m.t.Errorf("Expected call to ServiceMock.CreateChat with params: %#v", *m.CreateChatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateChat != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		m.t.Error("Expected call to ServiceMock.CreateChat")
	}
}

type mServiceMockDeleteChat struct {
	mock               *ServiceMock
	defaultExpectation *ServiceMockDeleteChatExpectation
	expectations       []*ServiceMockDeleteChatExpectation

	callArgs []*ServiceMockDeleteChatParams
	mutex    sync.RWMutex
}

// ServiceMockDeleteChatExpectation specifies expectation struct of the Service.DeleteChat
type ServiceMockDeleteChatExpectation struct {
	mock    *ServiceMock
	params  *ServiceMockDeleteChatParams
	results *ServiceMockDeleteChatResults
	Counter uint64
}

// ServiceMockDeleteChatParams contains parameters of the Service.DeleteChat
type ServiceMockDeleteChatParams struct {
	ctx context.Context
	id  int64
}

// ServiceMockDeleteChatResults contains results of the Service.DeleteChat
type ServiceMockDeleteChatResults struct {
	err error
}

// Expect sets up expected params for Service.DeleteChat
func (mmDeleteChat *mServiceMockDeleteChat) Expect(ctx context.Context, id int64) *mServiceMockDeleteChat {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ServiceMock.DeleteChat mock is already set by Set")
	}

	if mmDeleteChat.defaultExpectation == nil {
		mmDeleteChat.defaultExpectation = &ServiceMockDeleteChatExpectation{}
	}

	mmDeleteChat.defaultExpectation.params = &ServiceMockDeleteChatParams{ctx, id}
	for _, e := range mmDeleteChat.expectations {
		if minimock.Equal(e.params, mmDeleteChat.defaultExpectation.params) {
			mmDeleteChat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteChat.defaultExpectation.params)
		}
	}

	return mmDeleteChat
}

// Inspect accepts an inspector function that has same arguments as the Service.DeleteChat
func (mmDeleteChat *mServiceMockDeleteChat) Inspect(f func(ctx context.Context, id int64)) *mServiceMockDeleteChat {
	if mmDeleteChat.mock.inspectFuncDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("Inspect function is already set for ServiceMock.DeleteChat")
	}

	mmDeleteChat.mock.inspectFuncDeleteChat = f

	return mmDeleteChat
}

// Return sets up results that will be returned by Service.DeleteChat
func (mmDeleteChat *mServiceMockDeleteChat) Return(err error) *ServiceMock {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ServiceMock.DeleteChat mock is already set by Set")
	}

	if mmDeleteChat.defaultExpectation == nil {
		mmDeleteChat.defaultExpectation = &ServiceMockDeleteChatExpectation{mock: mmDeleteChat.mock}
	}
	mmDeleteChat.defaultExpectation.results = &ServiceMockDeleteChatResults{err}
	return mmDeleteChat.mock
}

// Set uses given function f to mock the Service.DeleteChat method
func (mmDeleteChat *mServiceMockDeleteChat) Set(f func(ctx context.Context, id int64) (err error)) *ServiceMock {
	if mmDeleteChat.defaultExpectation != nil {
		mmDeleteChat.mock.t.Fatalf("Default expectation is already set for the Service.DeleteChat method")
	}

	if len(mmDeleteChat.expectations) > 0 {
		mmDeleteChat.mock.t.Fatalf("Some expectations are already set for the Service.DeleteChat method")
	}

	mmDeleteChat.mock.funcDeleteChat = f
	return mmDeleteChat.mock
}

// When sets expectation for the Service.DeleteChat which will trigger the result defined by the following
// Then helper
func (mmDeleteChat *mServiceMockDeleteChat) When(ctx context.Context, id int64) *ServiceMockDeleteChatExpectation {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ServiceMock.DeleteChat mock is already set by Set")
	}

	expectation := &ServiceMockDeleteChatExpectation{
		mock:   mmDeleteChat.mock,
		params: &ServiceMockDeleteChatParams{ctx, id},
	}
	mmDeleteChat.expectations = append(mmDeleteChat.expectations, expectation)
	return expectation
}

// Then sets up Service.DeleteChat return parameters for the expectation previously defined by the When method
func (e *ServiceMockDeleteChatExpectation) Then(err error) *ServiceMock {
	e.results = &ServiceMockDeleteChatResults{err}
	return e.mock
}

// DeleteChat implements chat.Service
func (mmDeleteChat *ServiceMock) DeleteChat(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDeleteChat.beforeDeleteChatCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteChat.afterDeleteChatCounter, 1)

	if mmDeleteChat.inspectFuncDeleteChat != nil {
		mmDeleteChat.inspectFuncDeleteChat(ctx, id)
	}

	mm_params := ServiceMockDeleteChatParams{ctx, id}

	// Record call args
	mmDeleteChat.DeleteChatMock.mutex.Lock()
	mmDeleteChat.DeleteChatMock.callArgs = append(mmDeleteChat.DeleteChatMock.callArgs, &mm_params)
	mmDeleteChat.DeleteChatMock.mutex.Unlock()

	for _, e := range mmDeleteChat.DeleteChatMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeleteChat.DeleteChatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteChat.DeleteChatMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteChat.DeleteChatMock.defaultExpectation.params
		mm_got := ServiceMockDeleteChatParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteChat.t.Errorf("ServiceMock.DeleteChat got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteChat.DeleteChatMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteChat.t.Fatal("No results are set for the ServiceMock.DeleteChat")
		}
		return (*mm_results).err
	}
	if mmDeleteChat.funcDeleteChat != nil {
		return mmDeleteChat.funcDeleteChat(ctx, id)
	}
	mmDeleteChat.t.Fatalf("Unexpected call to ServiceMock.DeleteChat. %v %v", ctx, id)
	return
}

// DeleteChatAfterCounter returns a count of finished ServiceMock.DeleteChat invocations
func (mmDeleteChat *ServiceMock) DeleteChatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteChat.afterDeleteChatCounter)
}

// DeleteChatBeforeCounter returns a count of ServiceMock.DeleteChat invocations
func (mmDeleteChat *ServiceMock) DeleteChatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteChat.beforeDeleteChatCounter)
}

// Calls returns a list of arguments used in each call to ServiceMock.DeleteChat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteChat *mServiceMockDeleteChat) Calls() []*ServiceMockDeleteChatParams {
	mmDeleteChat.mutex.RLock()

	argCopy := make([]*ServiceMockDeleteChatParams, len(mmDeleteChat.callArgs))
	copy(argCopy, mmDeleteChat.callArgs)

	mmDeleteChat.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteChatDone returns true if the count of the DeleteChat invocations corresponds
// the number of defined expectations
func (m *ServiceMock) MinimockDeleteChatDone() bool {
	for _, e := range m.DeleteChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteChatCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteChat != nil && mm_atomic.LoadUint64(&m.afterDeleteChatCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteChatInspect logs each unmet expectation
func (m *ServiceMock) MinimockDeleteChatInspect() {
	for _, e := range m.DeleteChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ServiceMock.DeleteChat with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteChatCounter) < 1 {
		if m.DeleteChatMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ServiceMock.DeleteChat")
		} else {
			m.t.Errorf("Expected call to ServiceMock.DeleteChat with params: %#v", *m.DeleteChatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteChat != nil && mm_atomic.LoadUint64(&m.afterDeleteChatCounter) < 1 {
		m.t.Error("Expected call to ServiceMock.DeleteChat")
	}
}

type mServiceMockSendMessage struct {
	mock               *ServiceMock
	defaultExpectation *ServiceMockSendMessageExpectation
	expectations       []*ServiceMockSendMessageExpectation

	callArgs []*ServiceMockSendMessageParams
	mutex    sync.RWMutex
}

// ServiceMockSendMessageExpectation specifies expectation struct of the Service.SendMessage
type ServiceMockSendMessageExpectation struct {
	mock    *ServiceMock
	params  *ServiceMockSendMessageParams
	results *ServiceMockSendMessageResults
	Counter uint64
}

// ServiceMockSendMessageParams contains parameters of the Service.SendMessage
type ServiceMockSendMessageParams struct {
	ctx     context.Context
	message *modelMessage.Message
}

// ServiceMockSendMessageResults contains results of the Service.SendMessage
type ServiceMockSendMessageResults struct {
	err error
}

// Expect sets up expected params for Service.SendMessage
func (mmSendMessage *mServiceMockSendMessage) Expect(ctx context.Context, message *modelMessage.Message) *mServiceMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ServiceMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &ServiceMockSendMessageExpectation{}
	}

	mmSendMessage.defaultExpectation.params = &ServiceMockSendMessageParams{ctx, message}
	for _, e := range mmSendMessage.expectations {
		if minimock.Equal(e.params, mmSendMessage.defaultExpectation.params) {
			mmSendMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMessage.defaultExpectation.params)
		}
	}

	return mmSendMessage
}

// Inspect accepts an inspector function that has same arguments as the Service.SendMessage
func (mmSendMessage *mServiceMockSendMessage) Inspect(f func(ctx context.Context, message *modelMessage.Message)) *mServiceMockSendMessage {
	if mmSendMessage.mock.inspectFuncSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("Inspect function is already set for ServiceMock.SendMessage")
	}

	mmSendMessage.mock.inspectFuncSendMessage = f

	return mmSendMessage
}

// Return sets up results that will be returned by Service.SendMessage
func (mmSendMessage *mServiceMockSendMessage) Return(err error) *ServiceMock {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ServiceMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &ServiceMockSendMessageExpectation{mock: mmSendMessage.mock}
	}
	mmSendMessage.defaultExpectation.results = &ServiceMockSendMessageResults{err}
	return mmSendMessage.mock
}

// Set uses given function f to mock the Service.SendMessage method
func (mmSendMessage *mServiceMockSendMessage) Set(f func(ctx context.Context, message *modelMessage.Message) (err error)) *ServiceMock {
	if mmSendMessage.defaultExpectation != nil {
		mmSendMessage.mock.t.Fatalf("Default expectation is already set for the Service.SendMessage method")
	}

	if len(mmSendMessage.expectations) > 0 {
		mmSendMessage.mock.t.Fatalf("Some expectations are already set for the Service.SendMessage method")
	}

	mmSendMessage.mock.funcSendMessage = f
	return mmSendMessage.mock
}

// When sets expectation for the Service.SendMessage which will trigger the result defined by the following
// Then helper
func (mmSendMessage *mServiceMockSendMessage) When(ctx context.Context, message *modelMessage.Message) *ServiceMockSendMessageExpectation {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ServiceMock.SendMessage mock is already set by Set")
	}

	expectation := &ServiceMockSendMessageExpectation{
		mock:   mmSendMessage.mock,
		params: &ServiceMockSendMessageParams{ctx, message},
	}
	mmSendMessage.expectations = append(mmSendMessage.expectations, expectation)
	return expectation
}

// Then sets up Service.SendMessage return parameters for the expectation previously defined by the When method
func (e *ServiceMockSendMessageExpectation) Then(err error) *ServiceMock {
	e.results = &ServiceMockSendMessageResults{err}
	return e.mock
}

// SendMessage implements chat.Service
func (mmSendMessage *ServiceMock) SendMessage(ctx context.Context, message *modelMessage.Message) (err error) {
	mm_atomic.AddUint64(&mmSendMessage.beforeSendMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMessage.afterSendMessageCounter, 1)

	if mmSendMessage.inspectFuncSendMessage != nil {
		mmSendMessage.inspectFuncSendMessage(ctx, message)
	}

	mm_params := ServiceMockSendMessageParams{ctx, message}

	// Record call args
	mmSendMessage.SendMessageMock.mutex.Lock()
	mmSendMessage.SendMessageMock.callArgs = append(mmSendMessage.SendMessageMock.callArgs, &mm_params)
	mmSendMessage.SendMessageMock.mutex.Unlock()

	for _, e := range mmSendMessage.SendMessageMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendMessage.SendMessageMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendMessage.SendMessageMock.defaultExpectation.Counter, 1)
		mm_want := mmSendMessage.SendMessageMock.defaultExpectation.params
		mm_got := ServiceMockSendMessageParams{ctx, message}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMessage.t.Errorf("ServiceMock.SendMessage got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendMessage.SendMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmSendMessage.t.Fatal("No results are set for the ServiceMock.SendMessage")
		}
		return (*mm_results).err
	}
	if mmSendMessage.funcSendMessage != nil {
		return mmSendMessage.funcSendMessage(ctx, message)
	}
	mmSendMessage.t.Fatalf("Unexpected call to ServiceMock.SendMessage. %v %v", ctx, message)
	return
}

// SendMessageAfterCounter returns a count of finished ServiceMock.SendMessage invocations
func (mmSendMessage *ServiceMock) SendMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.afterSendMessageCounter)
}

// SendMessageBeforeCounter returns a count of ServiceMock.SendMessage invocations
func (mmSendMessage *ServiceMock) SendMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.beforeSendMessageCounter)
}

// Calls returns a list of arguments used in each call to ServiceMock.SendMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMessage *mServiceMockSendMessage) Calls() []*ServiceMockSendMessageParams {
	mmSendMessage.mutex.RLock()

	argCopy := make([]*ServiceMockSendMessageParams, len(mmSendMessage.callArgs))
	copy(argCopy, mmSendMessage.callArgs)

	mmSendMessage.mutex.RUnlock()

	return argCopy
}

// MinimockSendMessageDone returns true if the count of the SendMessage invocations corresponds
// the number of defined expectations
func (m *ServiceMock) MinimockSendMessageDone() bool {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendMessageInspect logs each unmet expectation
func (m *ServiceMock) MinimockSendMessageInspect() {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ServiceMock.SendMessage with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		if m.SendMessageMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ServiceMock.SendMessage")
		} else {
			m.t.Errorf("Expected call to ServiceMock.SendMessage with params: %#v", *m.SendMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		m.t.Error("Expected call to ServiceMock.SendMessage")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateChatInspect()

			m.MinimockDeleteChatInspect()

			m.MinimockSendMessageInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateChatDone() &&
		m.MinimockDeleteChatDone() &&
		m.MinimockSendMessageDone()
}
