// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// CloudWatchLogsQueryConverter accepts a `CloudWatchLogsQuery` object and generates the Go code to build this object using builders.
func CloudWatchLogsQueryConverter(input CloudWatchLogsQuery) string {
	calls := []string{
		`cloudwatch.NewCloudWatchLogsQueryBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`QueryMode(`)
		arg0 := cog.Dump(input.QueryMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Id != "" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Region != "" {

		buffer.WriteString(`Region(`)
		arg0 := fmt.Sprintf("%#v", input.Region)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Expression != nil && *input.Expression != "" {

		buffer.WriteString(`Expression(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Expression))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.StatsGroups != nil && len(input.StatsGroups) >= 1 {

		buffer.WriteString(`StatsGroups(`)
		tmparg0 := []string{}
		for _, arg1 := range input.StatsGroups {
			tmpstatsGroupsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpstatsGroupsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LogGroups != nil && len(input.LogGroups) >= 1 {

		buffer.WriteString(`LogGroups(`)
		tmparg0 := []string{}
		for _, arg1 := range input.LogGroups {
			tmplogGroupsarg1 := LogGroupConverter(arg1)
			tmparg0 = append(tmparg0, tmplogGroupsarg1)
		}
		arg0 := "[]cog.Builder[cloudwatch.LogGroup]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LogGroupNames != nil && len(input.LogGroupNames) >= 1 {

		buffer.WriteString(`LogGroupNames(`)
		tmparg0 := []string{}
		for _, arg1 := range input.LogGroupNames {
			tmplogGroupNamesarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmplogGroupNamesarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := cog.Dump(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
