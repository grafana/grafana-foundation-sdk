// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MetricStatConverter accepts a `MetricStat` object and generates the Go code to build this object using builders.
func MetricStatConverter(input MetricStat) string {
	calls := []string{
		`cloudwatch.NewMetricStatBuilder()`,
	}
	var buffer strings.Builder
	if input.Region != "" {

		buffer.WriteString(`Region(`)
		arg0 := fmt.Sprintf("%#v", input.Region)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Namespace != "" {

		buffer.WriteString(`Namespace(`)
		arg0 := fmt.Sprintf("%#v", input.Namespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MetricName != nil && *input.MetricName != "" {

		buffer.WriteString(`MetricName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MetricName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Dimensions != nil {

		buffer.WriteString(`Dimensions(`)
		arg0 := cog.Dump(*input.Dimensions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MatchExact != nil {

		buffer.WriteString(`MatchExact(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MatchExact))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Period != nil && *input.Period != "" {

		buffer.WriteString(`Period(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Period))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AccountId != nil && *input.AccountId != "" {

		buffer.WriteString(`AccountId(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AccountId))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Statistic != nil && *input.Statistic != "" {

		buffer.WriteString(`Statistic(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Statistic))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Statistics != nil && len(input.Statistics) >= 1 {

		buffer.WriteString(`Statistics(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Statistics {
			tmpstatisticsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpstatisticsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
