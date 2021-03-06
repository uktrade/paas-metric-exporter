// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/alphagov/paas-metric-exporter/metrics"
)

type FakeSender struct {
	GaugeStub        func(metric metrics.GaugeMetric) error
	gaugeMutex       sync.RWMutex
	gaugeArgsForCall []struct {
		metric metrics.GaugeMetric
	}
	gaugeReturns struct {
		result1 error
	}
	gaugeReturnsOnCall map[int]struct {
		result1 error
	}
	IncrStub        func(metric metrics.CounterMetric) error
	incrMutex       sync.RWMutex
	incrArgsForCall []struct {
		metric metrics.CounterMetric
	}
	incrReturns struct {
		result1 error
	}
	incrReturnsOnCall map[int]struct {
		result1 error
	}
	PrecisionTimingStub        func(metric metrics.PrecisionTimingMetric) error
	precisionTimingMutex       sync.RWMutex
	precisionTimingArgsForCall []struct {
		metric metrics.PrecisionTimingMetric
	}
	precisionTimingReturns struct {
		result1 error
	}
	precisionTimingReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSender) Gauge(metric metrics.GaugeMetric) error {
	fake.gaugeMutex.Lock()
	ret, specificReturn := fake.gaugeReturnsOnCall[len(fake.gaugeArgsForCall)]
	fake.gaugeArgsForCall = append(fake.gaugeArgsForCall, struct {
		metric metrics.GaugeMetric
	}{metric})
	fake.recordInvocation("Gauge", []interface{}{metric})
	fake.gaugeMutex.Unlock()
	if fake.GaugeStub != nil {
		return fake.GaugeStub(metric)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.gaugeReturns.result1
}

func (fake *FakeSender) GaugeCallCount() int {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return len(fake.gaugeArgsForCall)
}

func (fake *FakeSender) GaugeArgsForCall(i int) metrics.GaugeMetric {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return fake.gaugeArgsForCall[i].metric
}

func (fake *FakeSender) GaugeReturns(result1 error) {
	fake.GaugeStub = nil
	fake.gaugeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) GaugeReturnsOnCall(i int, result1 error) {
	fake.GaugeStub = nil
	if fake.gaugeReturnsOnCall == nil {
		fake.gaugeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.gaugeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) Incr(metric metrics.CounterMetric) error {
	fake.incrMutex.Lock()
	ret, specificReturn := fake.incrReturnsOnCall[len(fake.incrArgsForCall)]
	fake.incrArgsForCall = append(fake.incrArgsForCall, struct {
		metric metrics.CounterMetric
	}{metric})
	fake.recordInvocation("Incr", []interface{}{metric})
	fake.incrMutex.Unlock()
	if fake.IncrStub != nil {
		return fake.IncrStub(metric)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.incrReturns.result1
}

func (fake *FakeSender) IncrCallCount() int {
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	return len(fake.incrArgsForCall)
}

func (fake *FakeSender) IncrArgsForCall(i int) metrics.CounterMetric {
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	return fake.incrArgsForCall[i].metric
}

func (fake *FakeSender) IncrReturns(result1 error) {
	fake.IncrStub = nil
	fake.incrReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) IncrReturnsOnCall(i int, result1 error) {
	fake.IncrStub = nil
	if fake.incrReturnsOnCall == nil {
		fake.incrReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.incrReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
	fake.precisionTimingMutex.Lock()
	ret, specificReturn := fake.precisionTimingReturnsOnCall[len(fake.precisionTimingArgsForCall)]
	fake.precisionTimingArgsForCall = append(fake.precisionTimingArgsForCall, struct {
		metric metrics.PrecisionTimingMetric
	}{metric})
	fake.recordInvocation("PrecisionTiming", []interface{}{metric})
	fake.precisionTimingMutex.Unlock()
	if fake.PrecisionTimingStub != nil {
		return fake.PrecisionTimingStub(metric)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.precisionTimingReturns.result1
}

func (fake *FakeSender) PrecisionTimingCallCount() int {
	fake.precisionTimingMutex.RLock()
	defer fake.precisionTimingMutex.RUnlock()
	return len(fake.precisionTimingArgsForCall)
}

func (fake *FakeSender) PrecisionTimingArgsForCall(i int) metrics.PrecisionTimingMetric {
	fake.precisionTimingMutex.RLock()
	defer fake.precisionTimingMutex.RUnlock()
	return fake.precisionTimingArgsForCall[i].metric
}

func (fake *FakeSender) PrecisionTimingReturns(result1 error) {
	fake.PrecisionTimingStub = nil
	fake.precisionTimingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) PrecisionTimingReturnsOnCall(i int, result1 error) {
	fake.PrecisionTimingStub = nil
	if fake.precisionTimingReturnsOnCall == nil {
		fake.precisionTimingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.precisionTimingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	fake.precisionTimingMutex.RLock()
	defer fake.precisionTimingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSender) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.Sender = new(FakeSender)
