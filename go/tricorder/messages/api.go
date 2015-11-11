// Package messages provides the types needed to collect metrics via
// the go rpc calls or the REST API mentioned in the tricorder package.
package messages

import (
	"errors"
	"github.com/Symantec/tricorder/go/tricorder/types"
	"github.com/Symantec/tricorder/go/tricorder/units"
	"time"
)

var (
	// The MetricServer.GetMetric RPC call returns this if no
	// metric with given path exists.
	ErrMetricNotFound = errors.New("messages: No metric found.")
)

// RpcRangeWithCount represents the number of values within a
// particular range for go rpc
type RpcRangeWithCount struct {
	// Represents the lower bound of the range inclusive.
	// Ignore for the lowest range which never has a lower bound.
	Lower float64
	// Represents the upper bound of the range exclusive.
	// Ignore for the highest range which never has a upper bound.
	Upper float64
	// The number of values falling within the range.
	Count uint64
}

// RangeWithCount represents the number of values within a particular range
type RangeWithCount struct {
	// If non nil, represents the lower bound of the range inclusive.
	// nil means no lower bound
	Lower *float64 `json:"lower,omitempty"`
	// If non nil, represents the upper bound of the range exclusive.
	// nil means no upper bound
	Upper *float64 `json:"upper,omitempty"`
	// The number of values falling within the range.
	Count uint64 `json:"count"`
}

// RpcDistribution represents a distribution of values for go rpc.
type RpcDistribution struct {
	// The minimum value
	Min float64
	// The maximum value
	Max float64
	// The average value
	Average float64
	// The approximate median value
	Median float64
	// The sum
	Sum float64
	// The total number of values
	Count uint64
	// The number of values within each range
	Ranges []*RpcRangeWithCount
}

// Distribution represents a distribution of values.
type Distribution struct {
	// The minimum value
	Min float64 `json:"min"`
	// The maximum value
	Max float64 `json:"max"`
	// The average value
	Average float64 `json:"average"`
	// The approximate median value
	Median float64 `json:"median"`
	// The sum
	Sum float64 `json:"sum"`
	// The total number of values
	Count uint64 `json:"count"`
	// The number of values within each range
	Ranges []*RangeWithCount `json:"ranges,omitempty"`
}

// Duration represents a duration of time
// For negative durations, both Seconds and Nanoseconds are negative.
type Duration struct {
	Seconds     int64
	Nanoseconds int32
}

func NewDuration(d time.Duration) Duration {
	return newDuration(d)
}

// SinceEpoch returns the amount of time since unix epoch
func SinceEpoch(t time.Time) Duration {
	return sinceEpoch(t)
}

// AsGoDuration converts this duration to a go duration
func (d Duration) AsGoDuration() time.Duration {
	return d.asGoDuration()
}

// AsGoTime Converts this duration to a go time.
// This is the inverse of SinceEpoch.
func (d Duration) AsGoTime() time.Time {
	return d.asGoTime()
}

func (d Duration) String() string {
	return d._string()
}

// RpcValue represents the value of a metric for go rpc.
type RpcValue struct {
	// The value's type
	Kind types.Type
	// bool values stored here
	BoolValue bool
	// int values stored here
	IntValue int64
	// uint values stored here
	UintValue uint64
	// float values stored here
	FloatValue float64
	// string values are stored here.
	StringValue string
	// duration values stored here. Also time values are stored here
	// as time since Jan 1, 1970 GMT.
	DurationValue Duration
	// Distributions stored here
	DistributionValue *RpcDistribution
}

// Value represents the value of a metric.
type Value struct {
	// The value's type
	Kind types.Type `json:"kind"`
	// bool values stored here
	BoolValue *bool `json:"boolValue,omitempty"`
	// int values stored here
	IntValue *int64 `json:"intValue,omitempty"`
	// uint values stored here
	UintValue *uint64 `json:"uintValue,omitempty"`
	// float values stored here
	FloatValue *float64 `json:"floatValue,omitempty"`
	// string values are stored here. Also time values are stored here
	// as seconds after Jan 1, 1970 GMT in this format:
	// 1234567890.987654321
	StringValue *string `json:"stringValue,omitempty"`
	// Distributions stored here
	DistributionValue *Distribution `json:"distributionValue,omitempty"`
}

// RpcMetric represents a single metric for go rpc
type RpcMetric struct {
	// The absolute path to this metric
	Path string
	// The description of this metric
	Description string
	// The unit of measurement this metric represents
	Unit units.Unit
	// The value of this metric
	Value *RpcValue
}

// Metric represents a single metric
type Metric struct {
	// The absolute path to this metric
	Path string `json:"path"`
	// The description of this metric
	Description string `json:"description"`
	// The unit of measurement this metric represents
	Unit units.Unit `json:"unit"`
	// The value of this metric
	Value *Value `json:"value"`
}

// RpcMetricList represents a list of rpc metrics.
type RpcMetricList []*RpcMetric

// MetricList represents a list of metrics.
type MetricList []*Metric
