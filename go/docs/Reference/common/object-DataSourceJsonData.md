---
title: <span class="badge object-type-struct"></span> DataSourceJsonData
---
# <span class="badge object-type-struct"></span> DataSourceJsonData

TODO docs

## Definition

```go
type DataSourceJsonData struct {
    AuthType *string `json:"authType,omitempty"`
    DefaultRegion *string `json:"defaultRegion,omitempty"`
    Profile *string `json:"profile,omitempty"`
    ManageAlerts *bool `json:"manageAlerts,omitempty"`
    AlertmanagerUid *string `json:"alertmanagerUid,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataSourceJsonData` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dataSourceJsonData *DataSourceJsonData) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DataSourceJsonData` objects.

```go
func (dataSourceJsonData *DataSourceJsonData) Equals(other DataSourceJsonData) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DataSourceJsonData` fields for violations and returns them.

```go
func (dataSourceJsonData *DataSourceJsonData) Validate() error
```

## See also

 * <span class="badge builder"></span> [DataSourceJsonDataBuilder](./builder-DataSourceJsonDataBuilder.md)
