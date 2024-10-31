// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AnnotationQueryConverter accepts a `AnnotationQuery` object and generates the Go code to build this object using builders.
func AnnotationQueryConverter(input AnnotationQuery) string {
	calls := []string{
		`googlecloudmonitoring.NewAnnotationQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.ProjectName != "" {

		buffer.WriteString(`ProjectName(`)
		arg0 := fmt.Sprintf("%#v", input.ProjectName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CrossSeriesReducer != "" {

		buffer.WriteString(`CrossSeriesReducer(`)
		arg0 := fmt.Sprintf("%#v", input.CrossSeriesReducer)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AlignmentPeriod != nil && *input.AlignmentPeriod != "" {

		buffer.WriteString(`AlignmentPeriod(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AlignmentPeriod))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PerSeriesAligner != nil && *input.PerSeriesAligner != "" {

		buffer.WriteString(`PerSeriesAligner(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.PerSeriesAligner))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupBys != nil && len(input.GroupBys) >= 1 {

		buffer.WriteString(`GroupBys(`)
		tmparg0 := []string{}
		for _, arg1 := range input.GroupBys {
			tmpgroupBysarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpgroupBysarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filters != nil && len(input.Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Filters {
			tmpfiltersarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.View != nil && *input.View != "" {

		buffer.WriteString(`View(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.View))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SecondaryCrossSeriesReducer != nil && *input.SecondaryCrossSeriesReducer != "" {

		buffer.WriteString(`SecondaryCrossSeriesReducer(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SecondaryCrossSeriesReducer))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SecondaryAlignmentPeriod != nil && *input.SecondaryAlignmentPeriod != "" {

		buffer.WriteString(`SecondaryAlignmentPeriod(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SecondaryAlignmentPeriod))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SecondaryPerSeriesAligner != nil && *input.SecondaryPerSeriesAligner != "" {

		buffer.WriteString(`SecondaryPerSeriesAligner(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SecondaryPerSeriesAligner))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SecondaryGroupBys != nil && len(input.SecondaryGroupBys) >= 1 {

		buffer.WriteString(`SecondaryGroupBys(`)
		tmparg0 := []string{}
		for _, arg1 := range input.SecondaryGroupBys {
			tmpsecondaryGroupBysarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpsecondaryGroupBysarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Title != nil && *input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Title))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Preprocessor != nil {

		buffer.WriteString(`Preprocessor(`)
		arg0 := cog.Dump(*input.Preprocessor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Text != nil && *input.Text != "" {

		buffer.WriteString(`Text(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Text))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
