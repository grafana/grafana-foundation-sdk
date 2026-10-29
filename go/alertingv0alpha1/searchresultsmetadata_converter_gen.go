// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// SearchResultsMetadataConverter accepts a `SearchResultsMetadata` object and generates the Go code to build this object using builders.
func SearchResultsMetadataConverter(input SearchResultsMetadata) string {
	calls := []string{
		`alertingv0alpha1.NewSearchResultsMetadataBuilder()`,
	}
	var buffer strings.Builder
	if input.Continue != nil && *input.Continue != "" {

		buffer.WriteString(`Continue(`)
		arg0 := fmt.Sprintf("%#v", *input.Continue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TotalHits != nil {

		buffer.WriteString(`TotalHits(`)
		arg0 := fmt.Sprintf("%#v", *input.TotalHits)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TotalHitsRelation != nil {

		buffer.WriteString(`TotalHitsRelation(`)
		arg0 := cog.Dump(*input.TotalHitsRelation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
