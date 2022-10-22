// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/fallback
// gowrap: http://github.com/hexdigest/gowrap

package templatestests

//go:generate gowrap gen -p github.com/hexdigest/gowrap/templates_tests -i TestInterface -t ../templates/fallback -o interface_with_fallback.go -l ""

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// TestInterfaceWithFallback implements TestInterface interface wrapped with Prometheus metrics
type TestInterfaceWithFallback struct {
	implementations []TestInterface
	interval        time.Duration
}

// NewTestInterfaceWithFallback takes several implementations of the TestInterface and returns an instance of TestInterface
// which calls all implementations concurrently with given interval and returns first non-error response.
func NewTestInterfaceWithFallback(interval time.Duration, impls ...TestInterface) TestInterfaceWithFallback {
	return TestInterfaceWithFallback{implementations: impls, interval: interval}
}

// Channels implements TestInterface
func (_d TestInterfaceWithFallback) Channels(chA chan bool, chB chan<- bool, chanC <-chan bool) {
	type _resultStruct struct {
	}
	var _ch = make(chan _resultStruct, 0)

	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()

	go func() {
		for _i := 0; _i < len(_d.implementations); _i++ {
			go func(_impl TestInterface) {
				_impl.Channels(chA, chB, chanC)

				select {
				case _ch <- _resultStruct{}:
				default:
				}

			}(_d.implementations[_i])

			if _i < len(_d.implementations)-1 {
				<-_ticker.C
			}
		}
	}()

	<-_ch
	return

}

// ContextNoError implements TestInterface
func (_d TestInterfaceWithFallback) ContextNoError(ctx context.Context, a1 string, a2 string) {
	type _resultStruct struct {
	}
	var _ch = make(chan _resultStruct, 0)

	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	ctx, _cancelFunc := context.WithCancel(ctx)
	defer _cancelFunc()

	go func() {
		for _i := 0; _i < len(_d.implementations); _i++ {
			go func(_impl TestInterface) {
				_impl.ContextNoError(ctx, a1, a2)

				select {
				case _ch <- _resultStruct{}:
				case <-ctx.Done():
				}

			}(_d.implementations[_i])

			if _i < len(_d.implementations)-1 {
				<-_ticker.C
			}
		}
	}()

	for {
		select {
		case <-_ch:
			return
		case <-ctx.Done():

			return
		}
	}

}

// F implements TestInterface
func (_d TestInterfaceWithFallback) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	type _resultStruct struct {
		result1 string
		result2 string
		err     error
	}
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	ctx, _cancelFunc := context.WithCancel(ctx)
	defer _cancelFunc()

	go func() {
		for _i := 0; _i < len(_d.implementations); _i++ {
			go func(_impl TestInterface) {
				result1, result2, err := _impl.F(ctx, a1, a2...)
				if err != nil {
					err = fmt.Errorf("%T: %v", _impl, err)
				}

				select {
				case _ch <- _resultStruct{result1, result2, err}:
				case <-ctx.Done():
				}

			}(_d.implementations[_i])

			if _i < len(_d.implementations)-1 {
				<-_ticker.C
			}
		}
	}()

	for {
		select {
		case _res := <-_ch:
			if _res.err == nil {
				return _res.result1, _res.result2, _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
			if len(_errorsList) == len(_d.implementations) {
				err = fmt.Errorf(strings.Join(_errorsList, ";"))
				return
			}
		case <-ctx.Done():

			err = fmt.Errorf("%w: %s", ctx.Err(), strings.Join(_errorsList, ";"))

			return
		}
	}

}

// NoError implements TestInterface
func (_d TestInterfaceWithFallback) NoError(s1 string) (s2 string) {
	type _resultStruct struct {
		s2 string
	}
	var _ch = make(chan _resultStruct, 0)

	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()

	go func() {
		for _i := 0; _i < len(_d.implementations); _i++ {
			go func(_impl TestInterface) {
				s2 := _impl.NoError(s1)

				select {
				case _ch <- _resultStruct{s2}:
				default:
				}

			}(_d.implementations[_i])

			if _i < len(_d.implementations)-1 {
				<-_ticker.C
			}
		}
	}()

	_res := <-_ch
	return _res.s2

}

// NoParamsOrResults implements TestInterface
func (_d TestInterfaceWithFallback) NoParamsOrResults() {
	type _resultStruct struct {
	}
	var _ch = make(chan _resultStruct, 0)

	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()

	go func() {
		for _i := 0; _i < len(_d.implementations); _i++ {
			go func(_impl TestInterface) {
				_impl.NoParamsOrResults()

				select {
				case _ch <- _resultStruct{}:
				default:
				}

			}(_d.implementations[_i])

			if _i < len(_d.implementations)-1 {
				<-_ticker.C
			}
		}
	}()

	<-_ch
	return

}
