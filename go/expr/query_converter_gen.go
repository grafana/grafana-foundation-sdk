// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"fmt"
	"strings"

	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`expr.NewQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Version != "" && input.Version != "v0" {

		buffer.WriteString(`Version(`)
		arg0 := fmt.Sprintf("%#v", input.Version)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasourceConverter(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Expr).TypeMath != nil {

		buffer.WriteString(`TypeMath(`)
		arg0 := TypeMathConverter(*input.Spec.(*Expr).TypeMath)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Expr).TypeReduce != nil {

		buffer.WriteString(`TypeReduce(`)
		arg0 := TypeReduceConverter(*input.Spec.(*Expr).TypeReduce)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Expr).TypeResample != nil {

		buffer.WriteString(`TypeResample(`)
		arg0 := TypeResampleConverter(*input.Spec.(*Expr).TypeResample)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Expr).TypeClassicConditions != nil {

		buffer.WriteString(`TypeClassicConditions(`)
		arg0 := TypeClassicConditionsConverter(*input.Spec.(*Expr).TypeClassicConditions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Expr).TypeThreshold != nil {

		buffer.WriteString(`TypeThreshold(`)
		arg0 := TypeThresholdConverter(*input.Spec.(*Expr).TypeThreshold)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Expr).TypeSql != nil {

		buffer.WriteString(`TypeSql(`)
		arg0 := TypeSqlConverter(*input.Spec.(*Expr).TypeSql)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
