// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// QueryOptionsSpecConverter accepts a `QueryOptionsSpec` object and generates the Go code to build this object using builders.
func QueryOptionsSpecConverter(input QueryOptionsSpec) string {
	calls := []string{
		`dashboardv2beta1.NewQueryOptionsSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.TimeFrom != nil && *input.TimeFrom != "" {

		buffer.WriteString(`TimeFrom(`)
		arg0 := fmt.Sprintf("%#v", *input.TimeFrom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxDataPoints != nil {

		buffer.WriteString(`MaxDataPoints(`)
		arg0 := fmt.Sprintf("%#v", *input.MaxDataPoints)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeShift != nil && *input.TimeShift != "" {

		buffer.WriteString(`TimeShift(`)
		arg0 := fmt.Sprintf("%#v", *input.TimeShift)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryCachingTTL != nil {

		buffer.WriteString(`QueryCachingTTL(`)
		arg0 := fmt.Sprintf("%#v", *input.QueryCachingTTL)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Interval != nil && *input.Interval != "" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", *input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CacheTimeout != nil && *input.CacheTimeout != "" {

		buffer.WriteString(`CacheTimeout(`)
		arg0 := fmt.Sprintf("%#v", *input.CacheTimeout)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.HideTimeOverride != nil {

		buffer.WriteString(`HideTimeOverride(`)
		arg0 := fmt.Sprintf("%#v", *input.HideTimeOverride)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeCompare != nil && *input.TimeCompare != "" {

		buffer.WriteString(`TimeCompare(`)
		arg0 := fmt.Sprintf("%#v", *input.TimeCompare)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
