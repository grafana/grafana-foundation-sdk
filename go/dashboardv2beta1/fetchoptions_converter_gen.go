// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// FetchOptionsConverter accepts a `FetchOptions` object and generates the Go code to build this object using builders.
func FetchOptionsConverter(input FetchOptions) string {
	calls := []string{
		`dashboardv2beta1.NewFetchOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Method(`)
		arg0 := cog.Dump(input.Method)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Url != "" {

		buffer.WriteString(`Url(`)
		arg0 := fmt.Sprintf("%#v", input.Url)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Body != nil && *input.Body != "" {

		buffer.WriteString(`Body(`)
		arg0 := fmt.Sprintf("%#v", *input.Body)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryParams != nil && len(input.QueryParams) >= 1 {

		buffer.WriteString(`QueryParams(`)
		tmparg0 := []string{}
		for _, arg1 := range input.QueryParams {
			tmptmpqueryParamsarg1 := []string{}
			for _, arg1Value := range arg1 {
				tmparg1arg1Value := fmt.Sprintf("%#v", arg1Value)
				tmptmpqueryParamsarg1 = append(tmptmpqueryParamsarg1, tmparg1arg1Value)
			}
			tmpqueryParamsarg1 := "[]string{" + strings.Join(tmptmpqueryParamsarg1, ",\n") + "}"
			tmparg0 = append(tmparg0, tmpqueryParamsarg1)
		}
		arg0 := "[][]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Headers != nil && len(input.Headers) >= 1 {

		buffer.WriteString(`Headers(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Headers {
			tmptmpheadersarg1 := []string{}
			for _, arg1Value := range arg1 {
				tmparg1arg1Value := fmt.Sprintf("%#v", arg1Value)
				tmptmpheadersarg1 = append(tmptmpheadersarg1, tmparg1arg1Value)
			}
			tmpheadersarg1 := "[]string{" + strings.Join(tmptmpheadersarg1, ",\n") + "}"
			tmparg0 = append(tmparg0, tmpheadersarg1)
		}
		arg0 := "[][]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
