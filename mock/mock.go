// Package wait provides functionalities for mocking the wait components.
package wait

import "testing"

// Mocker creates a new Mocker.
type Mocker func(tb testing.TB) *Waiter

// NoMock creates a Mocker mock with no expectations.
var NoMock = Mock()

// Mock creates a Mocker mock with cleanup to ensure all the expectations are met.
func Mock(mocks ...func(e *Waiter)) Mocker {
	return func(tb testing.TB) *Waiter {
		tb.Helper()

		e := NewWaiter(tb)

		for _, m := range mocks {
			m(e)
		}

		return e
	}
}
