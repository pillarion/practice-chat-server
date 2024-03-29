// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mock

//go:generate minimock -i github.com/pillarion/practice-chat-server/internal/core/port/repository/journal.Repo -o repo_minimock.go -n RepoMock -p mock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	model "github.com/pillarion/practice-chat-server/internal/core/model/journal"
)

// RepoMock implements journal.Repo
type RepoMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcInsert          func(ctx context.Context, j *model.Journal) (i1 int64, err error)
	inspectFuncInsert   func(ctx context.Context, j *model.Journal)
	afterInsertCounter  uint64
	beforeInsertCounter uint64
	InsertMock          mRepoMockInsert
}

// NewRepoMock returns a mock for journal.Repo
func NewRepoMock(t minimock.Tester) *RepoMock {
	m := &RepoMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.InsertMock = mRepoMockInsert{mock: m}
	m.InsertMock.callArgs = []*RepoMockInsertParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRepoMockInsert struct {
	mock               *RepoMock
	defaultExpectation *RepoMockInsertExpectation
	expectations       []*RepoMockInsertExpectation

	callArgs []*RepoMockInsertParams
	mutex    sync.RWMutex
}

// RepoMockInsertExpectation specifies expectation struct of the Repo.Insert
type RepoMockInsertExpectation struct {
	mock    *RepoMock
	params  *RepoMockInsertParams
	results *RepoMockInsertResults
	Counter uint64
}

// RepoMockInsertParams contains parameters of the Repo.Insert
type RepoMockInsertParams struct {
	ctx context.Context
	j   *model.Journal
}

// RepoMockInsertResults contains results of the Repo.Insert
type RepoMockInsertResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for Repo.Insert
func (mmInsert *mRepoMockInsert) Expect(ctx context.Context, j *model.Journal) *mRepoMockInsert {
	if mmInsert.mock.funcInsert != nil {
		mmInsert.mock.t.Fatalf("RepoMock.Insert mock is already set by Set")
	}

	if mmInsert.defaultExpectation == nil {
		mmInsert.defaultExpectation = &RepoMockInsertExpectation{}
	}

	mmInsert.defaultExpectation.params = &RepoMockInsertParams{ctx, j}
	for _, e := range mmInsert.expectations {
		if minimock.Equal(e.params, mmInsert.defaultExpectation.params) {
			mmInsert.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInsert.defaultExpectation.params)
		}
	}

	return mmInsert
}

// Inspect accepts an inspector function that has same arguments as the Repo.Insert
func (mmInsert *mRepoMockInsert) Inspect(f func(ctx context.Context, j *model.Journal)) *mRepoMockInsert {
	if mmInsert.mock.inspectFuncInsert != nil {
		mmInsert.mock.t.Fatalf("Inspect function is already set for RepoMock.Insert")
	}

	mmInsert.mock.inspectFuncInsert = f

	return mmInsert
}

// Return sets up results that will be returned by Repo.Insert
func (mmInsert *mRepoMockInsert) Return(i1 int64, err error) *RepoMock {
	if mmInsert.mock.funcInsert != nil {
		mmInsert.mock.t.Fatalf("RepoMock.Insert mock is already set by Set")
	}

	if mmInsert.defaultExpectation == nil {
		mmInsert.defaultExpectation = &RepoMockInsertExpectation{mock: mmInsert.mock}
	}
	mmInsert.defaultExpectation.results = &RepoMockInsertResults{i1, err}
	return mmInsert.mock
}

// Set uses given function f to mock the Repo.Insert method
func (mmInsert *mRepoMockInsert) Set(f func(ctx context.Context, j *model.Journal) (i1 int64, err error)) *RepoMock {
	if mmInsert.defaultExpectation != nil {
		mmInsert.mock.t.Fatalf("Default expectation is already set for the Repo.Insert method")
	}

	if len(mmInsert.expectations) > 0 {
		mmInsert.mock.t.Fatalf("Some expectations are already set for the Repo.Insert method")
	}

	mmInsert.mock.funcInsert = f
	return mmInsert.mock
}

// When sets expectation for the Repo.Insert which will trigger the result defined by the following
// Then helper
func (mmInsert *mRepoMockInsert) When(ctx context.Context, j *model.Journal) *RepoMockInsertExpectation {
	if mmInsert.mock.funcInsert != nil {
		mmInsert.mock.t.Fatalf("RepoMock.Insert mock is already set by Set")
	}

	expectation := &RepoMockInsertExpectation{
		mock:   mmInsert.mock,
		params: &RepoMockInsertParams{ctx, j},
	}
	mmInsert.expectations = append(mmInsert.expectations, expectation)
	return expectation
}

// Then sets up Repo.Insert return parameters for the expectation previously defined by the When method
func (e *RepoMockInsertExpectation) Then(i1 int64, err error) *RepoMock {
	e.results = &RepoMockInsertResults{i1, err}
	return e.mock
}

// Insert implements journal.Repo
func (mmInsert *RepoMock) Insert(ctx context.Context, j *model.Journal) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmInsert.beforeInsertCounter, 1)
	defer mm_atomic.AddUint64(&mmInsert.afterInsertCounter, 1)

	if mmInsert.inspectFuncInsert != nil {
		mmInsert.inspectFuncInsert(ctx, j)
	}

	mm_params := RepoMockInsertParams{ctx, j}

	// Record call args
	mmInsert.InsertMock.mutex.Lock()
	mmInsert.InsertMock.callArgs = append(mmInsert.InsertMock.callArgs, &mm_params)
	mmInsert.InsertMock.mutex.Unlock()

	for _, e := range mmInsert.InsertMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmInsert.InsertMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInsert.InsertMock.defaultExpectation.Counter, 1)
		mm_want := mmInsert.InsertMock.defaultExpectation.params
		mm_got := RepoMockInsertParams{ctx, j}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInsert.t.Errorf("RepoMock.Insert got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInsert.InsertMock.defaultExpectation.results
		if mm_results == nil {
			mmInsert.t.Fatal("No results are set for the RepoMock.Insert")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmInsert.funcInsert != nil {
		return mmInsert.funcInsert(ctx, j)
	}
	mmInsert.t.Fatalf("Unexpected call to RepoMock.Insert. %v %v", ctx, j)
	return
}

// InsertAfterCounter returns a count of finished RepoMock.Insert invocations
func (mmInsert *RepoMock) InsertAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsert.afterInsertCounter)
}

// InsertBeforeCounter returns a count of RepoMock.Insert invocations
func (mmInsert *RepoMock) InsertBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsert.beforeInsertCounter)
}

// Calls returns a list of arguments used in each call to RepoMock.Insert.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInsert *mRepoMockInsert) Calls() []*RepoMockInsertParams {
	mmInsert.mutex.RLock()

	argCopy := make([]*RepoMockInsertParams, len(mmInsert.callArgs))
	copy(argCopy, mmInsert.callArgs)

	mmInsert.mutex.RUnlock()

	return argCopy
}

// MinimockInsertDone returns true if the count of the Insert invocations corresponds
// the number of defined expectations
func (m *RepoMock) MinimockInsertDone() bool {
	for _, e := range m.InsertMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsert != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertInspect logs each unmet expectation
func (m *RepoMock) MinimockInsertInspect() {
	for _, e := range m.InsertMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoMock.Insert with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		if m.InsertMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoMock.Insert")
		} else {
			m.t.Errorf("Expected call to RepoMock.Insert with params: %#v", *m.InsertMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsert != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		m.t.Error("Expected call to RepoMock.Insert")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepoMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockInsertInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepoMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockInsertDone()
}
