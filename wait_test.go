package wait_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/wait"
)

const (
	duration10 = 10 * time.Millisecond
	duration50 = 50 * time.Millisecond
)

func TestNoWait_Wait_ContextCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := wait.NoWait.Wait(ctx)

	assert.NoError(t, err)
}

func TestNoWait_Wait_ContextNotCanceled(t *testing.T) {
	t.Parallel()

	err := wait.NoWait.Wait(context.Background())

	assert.NoError(t, err)
}

func TestFunc_Wait(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario    string
		wait        wait.Func
		expectedErr error
	}{
		{
			scenario: "error",
			wait: func(context.Context) error {
				return errors.New("wait error")
			},
			expectedErr: errors.New("wait error"),
		},
		{
			scenario: "no error",
			wait: func(context.Context) error {
				return nil
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.wait.Wait(context.Background())

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestForDuration_Wait_ContextCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := wait.ForDuration(time.Minute).Wait(ctx)

	assert.Equal(t, context.Canceled, err)
}

func TestForDuration_Wait_ContextNotCanceled(t *testing.T) {
	t.Parallel()

	assertAfterDuration(t, duration10, func() {
		err := wait.ForDuration(duration10).Wait(context.Background())

		assert.NoError(t, err)
	})
}

func TestForSignal_Wait_ContextCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := wait.ForSignal(time.After(time.Minute)).Wait(ctx)

	assert.Equal(t, context.Canceled, err)
}

func TestForSignal_Wait_ContextNotCanceled(t *testing.T) {
	t.Parallel()

	assertAfterDuration(t, duration10, func() {
		err := wait.ForSignal(time.After(duration10)).Wait(context.Background())

		assert.NoError(t, err)
	})
}

func TestUntil_Wait_ContextCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := wait.Until(time.Now().Add(time.Minute)).Wait(ctx)

	assert.Equal(t, context.Canceled, err)
}

func TestUntil_Wait_ContextNotCanceled(t *testing.T) {
	t.Parallel()

	assertAfterDuration(t, duration10, func() {
		err := wait.Until(time.Now().Add(duration10)).Wait(context.Background())

		assert.NoError(t, err)
	})
}

func TestUntil_Wait_Past(t *testing.T) {
	t.Parallel()

	// When time is in the past, golang timer will fire immediately.
	// Therefore, we don't expect any error.

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := wait.Until(time.Now().Add(-time.Hour)).Wait(ctx)

	assert.NoError(t, err)
}

func TestUntil_Wait_Future(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	assertAfterDuration(t, duration50, func() {
		err := wait.Until(time.Now().Add(duration50)).Wait(ctx)

		assert.NoError(t, err)
	})
}

func assertAfterDuration(t *testing.T, expected time.Duration, f func()) {
	t.Helper()

	startTime := time.Now()

	f()

	endTime := time.Now()

	assert.LessOrEqualf(t, startTime.Add(expected), endTime, "expected to be after %s, but was %s", expected, endTime.Sub(startTime))
}
