// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"
)

// SearchResultsConverter accepts a `SearchResults` object and generates the Go code to build this object using builders.
func SearchResultsConverter(input SearchResults) string {
	calls := []string{
		`alertingv0alpha1.NewSearchResultsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Metadata(`)
		arg0 := SearchResultsMetadataConverter(input.Metadata)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Items != nil && len(input.Items) >= 1 {

		buffer.WriteString(`Items(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Items {
			tmpitemsarg1 := SearchResultHitConverter(arg1)
			tmparg0 = append(tmparg0, tmpitemsarg1)
		}
		arg0 := "[]cog.Builder[alertingv0alpha1.SearchResultHit]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Facets != nil {

		buffer.WriteString(`Facets(`)
		arg0 := "map[string][]cog.Builder[alertingv0alpha1.FacetValue]{"
		for key, arg1 := range input.Facets {
			tmptmpfacetsarg1 := []string{}
			for _, arg1Value := range arg1 {
				tmparg1arg1Value := FacetValueConverter(arg1Value)
				tmptmpfacetsarg1 = append(tmptmpfacetsarg1, tmparg1arg1Value)
			}
			tmpfacetsarg1 := "[]cog.Builder[alertingv0alpha1.FacetValue]{" + strings.Join(tmptmpfacetsarg1, ",\n") + "}"
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpfacetsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
