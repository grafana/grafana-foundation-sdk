// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	"fmt"
	"strings"
)

// PrometheusDataqueryScopeConverter accepts a `PrometheusDataqueryScope` object and generates the Go code to build this object using builders.
func PrometheusDataqueryScopeConverter(input PrometheusDataqueryScope) string {
	calls := []string{
		`prometheus.NewPrometheusDataqueryScopeBuilder()`,
	}
	var buffer strings.Builder
	if input.Matchers != "" {

		buffer.WriteString(`Matchers(`)
		arg0 := fmt.Sprintf("%#v", input.Matchers)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
