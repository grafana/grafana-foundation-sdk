// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ResourceGroupsQueryConverter accepts a `ResourceGroupsQuery` object and generates the Go code to build this object using builders.
func ResourceGroupsQueryConverter(input ResourceGroupsQuery) string {
	calls := []string{
		`azuremonitor.NewResourceGroupsQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RawQuery != nil && *input.RawQuery != "" {

		buffer.WriteString(`RawQuery(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.RawQuery))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Subscription != "" {

		buffer.WriteString(`Subscription(`)
		arg0 := fmt.Sprintf("%#v", input.Subscription)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
