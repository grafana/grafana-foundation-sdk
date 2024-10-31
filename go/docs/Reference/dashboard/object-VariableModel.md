---
title: <span class="badge object-type-struct"></span> VariableModel
---
# <span class="badge object-type-struct"></span> VariableModel

A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.

## Definition

```go
type VariableModel struct {
    // Unique numeric identifier for the variable.
    Id string `json:"id"`
    // Type of variable
    Type dashboard.VariableType `json:"type"`
    // Name of variable
    Name string `json:"name"`
    // Optional display name
    Label *string `json:"label,omitempty"`
    // Visibility configuration for the variable
    Hide dashboard.VariableHide `json:"hide"`
    // Whether the variable value should be managed by URL query params or not
    SkipUrlSync bool `json:"skipUrlSync"`
    // Description of variable. It can be defined but `null`.
    Description *string `json:"description,omitempty"`
    // Query used to fetch values for a variable
    Query *dashboard.StringOrMap `json:"query,omitempty"`
    // Data source used to fetch values for a variable. It can be defined but `null`.
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
    // Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
    AllFormat *string `json:"allFormat,omitempty"`
    // Shows current selected variable text/value on the dashboard
    Current *dashboard.VariableOption `json:"current,omitempty"`
    // Whether multiple values can be selected or not from variable value list
    Multi *bool `json:"multi,omitempty"`
    // Options that can be selected for a variable.
    Options []dashboard.VariableOption `json:"options,omitempty"`
    // Options to config when to refresh a variable
    Refresh *dashboard.VariableRefresh `json:"refresh,omitempty"`
    // Options sort order
    Sort *dashboard.VariableSort `json:"sort,omitempty"`
    // Whether all value option is available or not
    IncludeAll *bool `json:"includeAll,omitempty"`
    // Custom all value
    AllValue *string `json:"allValue,omitempty"`
    // Optional field, if you want to extract part of a series name or metric node segment.
    // Named capture groups can be used to separate the display text and value.
    Regex *string `json:"regex,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VariableModel` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (variableModel *VariableModel) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `VariableModel` objects.

```go
func (variableModel *VariableModel) Equals(other VariableModel) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VariableModel` fields for violations and returns them.

```go
func (variableModel *VariableModel) Validate() error
```

## See also

 * <span class="badge builder"></span> [AdHocVariableBuilder](./builder-AdHocVariableBuilder.md)
 * <span class="badge builder"></span> [ConstantVariableBuilder](./builder-ConstantVariableBuilder.md)
 * <span class="badge builder"></span> [CustomVariableBuilder](./builder-CustomVariableBuilder.md)
 * <span class="badge builder"></span> [DatasourceVariableBuilder](./builder-DatasourceVariableBuilder.md)
 * <span class="badge builder"></span> [IntervalVariableBuilder](./builder-IntervalVariableBuilder.md)
 * <span class="badge builder"></span> [QueryVariableBuilder](./builder-QueryVariableBuilder.md)
 * <span class="badge builder"></span> [TextBoxVariableBuilder](./builder-TextBoxVariableBuilder.md)
