// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type QueryFormat int64

const (
	QueryFormatTimeseries QueryFormat = 0
	QueryFormatTable      QueryFormat = 1
)

type QueryPriority string

const (
	QueryPriorityInteractive QueryPriority = "INTERACTIVE"
	QueryPriorityBatch       QueryPriority = "BATCH"
)

type EditorMode string

const (
	EditorModeCode    EditorMode = "code"
	EditorModeBuilder EditorMode = "builder"
)

type SQLExpression struct {
	Columns []QueryEditorFunctionExpression `json:"columns,omitempty"`
	From    *string                         `json:"from,omitempty"`
	// whereJsonTree?: _
	WhereString      *string                        `json:"whereString,omitempty"`
	GroupBy          []QueryEditorGroupByExpression `json:"groupBy,omitempty"`
	OrderBy          *QueryEditorPropertyExpression `json:"orderBy,omitempty"`
	OrderByDirection *OrderByDirection              `json:"orderByDirection,omitempty"`
	Limit            *int64                         `json:"limit,omitempty"`
	Offset           *int64                         `json:"offset,omitempty"`
}

// NewSQLExpression creates a new SQLExpression object.
func NewSQLExpression() *SQLExpression {
	return &SQLExpression{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SQLExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SQLExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "columns"
	if fields["columns"] != nil {
		if string(fields["columns"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["columns"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 QueryEditorFunctionExpression

				result1 = QueryEditorFunctionExpression{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("columns["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Columns = append(resource.Columns, result1)
			}

		}
		delete(fields, "columns")

	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}

		}
		delete(fields, "from")

	}
	// Field "whereString"
	if fields["whereString"] != nil {
		if string(fields["whereString"]) != "null" {
			if err := json.Unmarshal(fields["whereString"], &resource.WhereString); err != nil {
				errs = append(errs, cog.MakeBuildErrors("whereString", err)...)
			}

		}
		delete(fields, "whereString")

	}
	// Field "groupBy"
	if fields["groupBy"] != nil {
		if string(fields["groupBy"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["groupBy"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 QueryEditorGroupByExpression

				result1 = QueryEditorGroupByExpression{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("groupBy["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.GroupBy = append(resource.GroupBy, result1)
			}

		}
		delete(fields, "groupBy")

	}
	// Field "orderBy"
	if fields["orderBy"] != nil {
		if string(fields["orderBy"]) != "null" {

			resource.OrderBy = &QueryEditorPropertyExpression{}
			if err := resource.OrderBy.UnmarshalJSONStrict(fields["orderBy"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orderBy", err)...)
			}

		}
		delete(fields, "orderBy")

	}
	// Field "orderByDirection"
	if fields["orderByDirection"] != nil {
		if string(fields["orderByDirection"]) != "null" {
			if err := json.Unmarshal(fields["orderByDirection"], &resource.OrderByDirection); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orderByDirection", err)...)
			}

		}
		delete(fields, "orderByDirection")

	}
	// Field "limit"
	if fields["limit"] != nil {
		if string(fields["limit"]) != "null" {
			if err := json.Unmarshal(fields["limit"], &resource.Limit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("limit", err)...)
			}

		}
		delete(fields, "limit")

	}
	// Field "offset"
	if fields["offset"] != nil {
		if string(fields["offset"]) != "null" {
			if err := json.Unmarshal(fields["offset"], &resource.Offset); err != nil {
				errs = append(errs, cog.MakeBuildErrors("offset", err)...)
			}

		}
		delete(fields, "offset")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SQLExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SQLExpression` objects.
func (resource SQLExpression) Equals(other SQLExpression) bool {

	if len(resource.Columns) != len(other.Columns) {
		return false
	}

	for i1 := range resource.Columns {
		if !resource.Columns[i1].Equals(other.Columns[i1]) {
			return false
		}
	}
	if resource.From == nil && other.From != nil || resource.From != nil && other.From == nil {
		return false
	}

	if resource.From != nil {
		if *resource.From != *other.From {
			return false
		}
	}
	if resource.WhereString == nil && other.WhereString != nil || resource.WhereString != nil && other.WhereString == nil {
		return false
	}

	if resource.WhereString != nil {
		if *resource.WhereString != *other.WhereString {
			return false
		}
	}

	if len(resource.GroupBy) != len(other.GroupBy) {
		return false
	}

	for i1 := range resource.GroupBy {
		if !resource.GroupBy[i1].Equals(other.GroupBy[i1]) {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if !resource.OrderBy.Equals(*other.OrderBy) {
			return false
		}
	}
	if resource.OrderByDirection == nil && other.OrderByDirection != nil || resource.OrderByDirection != nil && other.OrderByDirection == nil {
		return false
	}

	if resource.OrderByDirection != nil {
		if *resource.OrderByDirection != *other.OrderByDirection {
			return false
		}
	}
	if resource.Limit == nil && other.Limit != nil || resource.Limit != nil && other.Limit == nil {
		return false
	}

	if resource.Limit != nil {
		if *resource.Limit != *other.Limit {
			return false
		}
	}
	if resource.Offset == nil && other.Offset != nil || resource.Offset != nil && other.Offset == nil {
		return false
	}

	if resource.Offset != nil {
		if *resource.Offset != *other.Offset {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SQLExpression` fields for violations and returns them.
func (resource SQLExpression) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Columns {
		if err := resource.Columns[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("columns["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.GroupBy {
		if err := resource.GroupBy[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("groupBy["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.OrderBy != nil {
		if err := resource.OrderBy.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("orderBy", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorExpressionType string

const (
	QueryEditorExpressionTypeProperty          QueryEditorExpressionType = "property"
	QueryEditorExpressionTypeOperator          QueryEditorExpressionType = "operator"
	QueryEditorExpressionTypeOr                QueryEditorExpressionType = "or"
	QueryEditorExpressionTypeAnd               QueryEditorExpressionType = "and"
	QueryEditorExpressionTypeGroupBy           QueryEditorExpressionType = "groupBy"
	QueryEditorExpressionTypeFunction          QueryEditorExpressionType = "function"
	QueryEditorExpressionTypeFunctionParameter QueryEditorExpressionType = "functionParameter"
)

type QueryEditorFunctionExpression struct {
	Type       string                                   `json:"type"`
	Name       *string                                  `json:"name,omitempty"`
	Parameters []QueryEditorFunctionParameterExpression `json:"parameters,omitempty"`
}

// NewQueryEditorFunctionExpression creates a new QueryEditorFunctionExpression object.
func NewQueryEditorFunctionExpression() *QueryEditorFunctionExpression {
	return &QueryEditorFunctionExpression{
		Type: "function",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorFunctionExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorFunctionExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "parameters"
	if fields["parameters"] != nil {
		if string(fields["parameters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["parameters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 QueryEditorFunctionParameterExpression

				result1 = QueryEditorFunctionParameterExpression{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("parameters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Parameters = append(resource.Parameters, result1)
			}

		}
		delete(fields, "parameters")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorFunctionExpression` objects.
func (resource QueryEditorFunctionExpression) Equals(other QueryEditorFunctionExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	if len(resource.Parameters) != len(other.Parameters) {
		return false
	}

	for i1 := range resource.Parameters {
		if !resource.Parameters[i1].Equals(other.Parameters[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorFunctionExpression` fields for violations and returns them.
func (resource QueryEditorFunctionExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "function") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == function"),
		)...)
	}

	for i1 := range resource.Parameters {
		if err := resource.Parameters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("parameters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorFunctionParameterExpression struct {
	Type string  `json:"type"`
	Name *string `json:"name,omitempty"`
}

// NewQueryEditorFunctionParameterExpression creates a new QueryEditorFunctionParameterExpression object.
func NewQueryEditorFunctionParameterExpression() *QueryEditorFunctionParameterExpression {
	return &QueryEditorFunctionParameterExpression{
		Type: "functionParameter",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorFunctionParameterExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorFunctionParameterExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionParameterExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorFunctionParameterExpression` objects.
func (resource QueryEditorFunctionParameterExpression) Equals(other QueryEditorFunctionParameterExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorFunctionParameterExpression` fields for violations and returns them.
func (resource QueryEditorFunctionParameterExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "functionParameter") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == functionParameter"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorGroupByExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
}

// NewQueryEditorGroupByExpression creates a new QueryEditorGroupByExpression object.
func NewQueryEditorGroupByExpression() *QueryEditorGroupByExpression {
	return &QueryEditorGroupByExpression{
		Type:     "groupBy",
		Property: *NewQueryEditorProperty(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorGroupByExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorGroupByExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "property"
	if fields["property"] != nil {
		if string(fields["property"]) != "null" {

			resource.Property = QueryEditorProperty{}
			if err := resource.Property.UnmarshalJSONStrict(fields["property"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("property", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is null"))...)

		}
		delete(fields, "property")
	} else {
		errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorGroupByExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorGroupByExpression` objects.
func (resource QueryEditorGroupByExpression) Equals(other QueryEditorGroupByExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Property.Equals(other.Property) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorGroupByExpression` fields for violations and returns them.
func (resource QueryEditorGroupByExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "groupBy") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == groupBy"),
		)...)
	}
	if err := resource.Property.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("property", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorProperty struct {
	Type QueryEditorPropertyType `json:"type"`
	Name *string                 `json:"name,omitempty"`
}

// NewQueryEditorProperty creates a new QueryEditorProperty object.
func NewQueryEditorProperty() *QueryEditorProperty {
	return &QueryEditorProperty{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorProperty` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorProperty) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorProperty", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorProperty` objects.
func (resource QueryEditorProperty) Equals(other QueryEditorProperty) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorProperty` fields for violations and returns them.
func (resource QueryEditorProperty) Validate() error {
	return nil
}

type QueryEditorPropertyType string

const (
	QueryEditorPropertyTypeString QueryEditorPropertyType = "string"
)

type QueryEditorPropertyExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
}

// NewQueryEditorPropertyExpression creates a new QueryEditorPropertyExpression object.
func NewQueryEditorPropertyExpression() *QueryEditorPropertyExpression {
	return &QueryEditorPropertyExpression{
		Type:     "property",
		Property: *NewQueryEditorProperty(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorPropertyExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorPropertyExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "property"
	if fields["property"] != nil {
		if string(fields["property"]) != "null" {

			resource.Property = QueryEditorProperty{}
			if err := resource.Property.UnmarshalJSONStrict(fields["property"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("property", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is null"))...)

		}
		delete(fields, "property")
	} else {
		errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorPropertyExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorPropertyExpression` objects.
func (resource QueryEditorPropertyExpression) Equals(other QueryEditorPropertyExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Property.Equals(other.Property) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorPropertyExpression` fields for violations and returns them.
func (resource QueryEditorPropertyExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "property") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == property"),
		)...)
	}
	if err := resource.Property.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("property", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type OrderByDirection string

const (
	OrderByDirectionASC  OrderByDirection = "ASC"
	OrderByDirectionDESC OrderByDirection = "DESC"
)

// Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
type Dataquery struct {
	Dataset          *string        `json:"dataset,omitempty"`
	Table            *string        `json:"table,omitempty"`
	Project          *string        `json:"project,omitempty"`
	Format           QueryFormat    `json:"format"`
	RawQuery         *bool          `json:"rawQuery,omitempty"`
	RawSql           string         `json:"rawSql"`
	Location         *string        `json:"location,omitempty"`
	Partitioned      *bool          `json:"partitioned,omitempty"`
	PartitionedField *string        `json:"partitionedField,omitempty"`
	ConvertToUTC     *bool          `json:"convertToUTC,omitempty"`
	Sharded          *bool          `json:"sharded,omitempty"`
	QueryPriority    *QueryPriority `json:"queryPriority,omitempty"`
	TimeShift        *string        `json:"timeShift,omitempty"`
	EditorMode       *EditorMode    `json:"editorMode,omitempty"`
	Sql              *SQLExpression `json:"sql,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return "grafana-bigquery-datasource"
}

// NewDataquery creates a new Dataquery object.
func NewDataquery() *Dataquery {
	return &Dataquery{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dataquery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dataquery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "dataset"
	if fields["dataset"] != nil {
		if string(fields["dataset"]) != "null" {
			if err := json.Unmarshal(fields["dataset"], &resource.Dataset); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dataset", err)...)
			}

		}
		delete(fields, "dataset")

	}
	// Field "table"
	if fields["table"] != nil {
		if string(fields["table"]) != "null" {
			if err := json.Unmarshal(fields["table"], &resource.Table); err != nil {
				errs = append(errs, cog.MakeBuildErrors("table", err)...)
			}

		}
		delete(fields, "table")

	}
	// Field "project"
	if fields["project"] != nil {
		if string(fields["project"]) != "null" {
			if err := json.Unmarshal(fields["project"], &resource.Project); err != nil {
				errs = append(errs, cog.MakeBuildErrors("project", err)...)
			}

		}
		delete(fields, "project")

	}
	// Field "format"
	if fields["format"] != nil {
		if string(fields["format"]) != "null" {
			if err := json.Unmarshal(fields["format"], &resource.Format); err != nil {
				errs = append(errs, cog.MakeBuildErrors("format", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("format", errors.New("required field is null"))...)

		}
		delete(fields, "format")
	} else {
		errs = append(errs, cog.MakeBuildErrors("format", errors.New("required field is missing from input"))...)
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "rawSql"
	if fields["rawSql"] != nil {
		if string(fields["rawSql"]) != "null" {
			if err := json.Unmarshal(fields["rawSql"], &resource.RawSql); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawSql", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("rawSql", errors.New("required field is null"))...)

		}
		delete(fields, "rawSql")
	} else {
		errs = append(errs, cog.MakeBuildErrors("rawSql", errors.New("required field is missing from input"))...)
	}
	// Field "location"
	if fields["location"] != nil {
		if string(fields["location"]) != "null" {
			if err := json.Unmarshal(fields["location"], &resource.Location); err != nil {
				errs = append(errs, cog.MakeBuildErrors("location", err)...)
			}

		}
		delete(fields, "location")

	}
	// Field "partitioned"
	if fields["partitioned"] != nil {
		if string(fields["partitioned"]) != "null" {
			if err := json.Unmarshal(fields["partitioned"], &resource.Partitioned); err != nil {
				errs = append(errs, cog.MakeBuildErrors("partitioned", err)...)
			}

		}
		delete(fields, "partitioned")

	}
	// Field "partitionedField"
	if fields["partitionedField"] != nil {
		if string(fields["partitionedField"]) != "null" {
			if err := json.Unmarshal(fields["partitionedField"], &resource.PartitionedField); err != nil {
				errs = append(errs, cog.MakeBuildErrors("partitionedField", err)...)
			}

		}
		delete(fields, "partitionedField")

	}
	// Field "convertToUTC"
	if fields["convertToUTC"] != nil {
		if string(fields["convertToUTC"]) != "null" {
			if err := json.Unmarshal(fields["convertToUTC"], &resource.ConvertToUTC); err != nil {
				errs = append(errs, cog.MakeBuildErrors("convertToUTC", err)...)
			}

		}
		delete(fields, "convertToUTC")

	}
	// Field "sharded"
	if fields["sharded"] != nil {
		if string(fields["sharded"]) != "null" {
			if err := json.Unmarshal(fields["sharded"], &resource.Sharded); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sharded", err)...)
			}

		}
		delete(fields, "sharded")

	}
	// Field "queryPriority"
	if fields["queryPriority"] != nil {
		if string(fields["queryPriority"]) != "null" {
			if err := json.Unmarshal(fields["queryPriority"], &resource.QueryPriority); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryPriority", err)...)
			}

		}
		delete(fields, "queryPriority")

	}
	// Field "timeShift"
	if fields["timeShift"] != nil {
		if string(fields["timeShift"]) != "null" {
			if err := json.Unmarshal(fields["timeShift"], &resource.TimeShift); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeShift", err)...)
			}

		}
		delete(fields, "timeShift")

	}
	// Field "editorMode"
	if fields["editorMode"] != nil {
		if string(fields["editorMode"]) != "null" {
			if err := json.Unmarshal(fields["editorMode"], &resource.EditorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("editorMode", err)...)
			}

		}
		delete(fields, "editorMode")

	}
	// Field "sql"
	if fields["sql"] != nil {
		if string(fields["sql"]) != "null" {

			resource.Sql = &SQLExpression{}
			if err := resource.Sql.UnmarshalJSONStrict(fields["sql"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sql", err)...)
			}

		}
		delete(fields, "sql")

	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dataquery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
func (resource Dataquery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(Dataquery)
	if !ok {
		return false
	}
	if resource.Dataset == nil && other.Dataset != nil || resource.Dataset != nil && other.Dataset == nil {
		return false
	}

	if resource.Dataset != nil {
		if *resource.Dataset != *other.Dataset {
			return false
		}
	}
	if resource.Table == nil && other.Table != nil || resource.Table != nil && other.Table == nil {
		return false
	}

	if resource.Table != nil {
		if *resource.Table != *other.Table {
			return false
		}
	}
	if resource.Project == nil && other.Project != nil || resource.Project != nil && other.Project == nil {
		return false
	}

	if resource.Project != nil {
		if *resource.Project != *other.Project {
			return false
		}
	}
	if resource.Format != other.Format {
		return false
	}
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.RawSql != other.RawSql {
		return false
	}
	if resource.Location == nil && other.Location != nil || resource.Location != nil && other.Location == nil {
		return false
	}

	if resource.Location != nil {
		if *resource.Location != *other.Location {
			return false
		}
	}
	if resource.Partitioned == nil && other.Partitioned != nil || resource.Partitioned != nil && other.Partitioned == nil {
		return false
	}

	if resource.Partitioned != nil {
		if *resource.Partitioned != *other.Partitioned {
			return false
		}
	}
	if resource.PartitionedField == nil && other.PartitionedField != nil || resource.PartitionedField != nil && other.PartitionedField == nil {
		return false
	}

	if resource.PartitionedField != nil {
		if *resource.PartitionedField != *other.PartitionedField {
			return false
		}
	}
	if resource.ConvertToUTC == nil && other.ConvertToUTC != nil || resource.ConvertToUTC != nil && other.ConvertToUTC == nil {
		return false
	}

	if resource.ConvertToUTC != nil {
		if *resource.ConvertToUTC != *other.ConvertToUTC {
			return false
		}
	}
	if resource.Sharded == nil && other.Sharded != nil || resource.Sharded != nil && other.Sharded == nil {
		return false
	}

	if resource.Sharded != nil {
		if *resource.Sharded != *other.Sharded {
			return false
		}
	}
	if resource.QueryPriority == nil && other.QueryPriority != nil || resource.QueryPriority != nil && other.QueryPriority == nil {
		return false
	}

	if resource.QueryPriority != nil {
		if *resource.QueryPriority != *other.QueryPriority {
			return false
		}
	}
	if resource.TimeShift == nil && other.TimeShift != nil || resource.TimeShift != nil && other.TimeShift == nil {
		return false
	}

	if resource.TimeShift != nil {
		if *resource.TimeShift != *other.TimeShift {
			return false
		}
	}
	if resource.EditorMode == nil && other.EditorMode != nil || resource.EditorMode != nil && other.EditorMode == nil {
		return false
	}

	if resource.EditorMode != nil {
		if *resource.EditorMode != *other.EditorMode {
			return false
		}
	}
	if resource.Sql == nil && other.Sql != nil || resource.Sql != nil && other.Sql == nil {
		return false
	}

	if resource.Sql != nil {
		if !resource.Sql.Equals(*other.Sql) {
			return false
		}
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.
func (resource Dataquery) Validate() error {
	var errs cog.BuildErrors
	if resource.Sql != nil {
		if err := resource.Sql.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("sql", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to grafana-bigquery-datasource dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "grafana-bigquery-datasource",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery Dataquery
			if cast, ok := input.(*Dataquery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(Dataquery)
			}
			return DataqueryConverter(dataquery)
		},
	}
}
