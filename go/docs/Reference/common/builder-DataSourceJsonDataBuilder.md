---
title: <span class="badge builder"></span> DataSourceJsonDataBuilder
---
# <span class="badge builder"></span> DataSourceJsonDataBuilder

## Constructor

```go
func NewDataSourceJsonDataBuilder() *DataSourceJsonDataBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataSourceJsonDataBuilder) Build() (DataSourceJsonData, error)
```

### <span class="badge object-method"></span> AlertmanagerUid

```go
func (builder *DataSourceJsonDataBuilder) AlertmanagerUid(alertmanagerUid string) *DataSourceJsonDataBuilder
```

### <span class="badge object-method"></span> AuthType

```go
func (builder *DataSourceJsonDataBuilder) AuthType(authType string) *DataSourceJsonDataBuilder
```

### <span class="badge object-method"></span> DefaultRegion

```go
func (builder *DataSourceJsonDataBuilder) DefaultRegion(defaultRegion string) *DataSourceJsonDataBuilder
```

### <span class="badge object-method"></span> ManageAlerts

```go
func (builder *DataSourceJsonDataBuilder) ManageAlerts(manageAlerts bool) *DataSourceJsonDataBuilder
```

### <span class="badge object-method"></span> Profile

```go
func (builder *DataSourceJsonDataBuilder) Profile(profile string) *DataSourceJsonDataBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DataSourceJsonData](./object-DataSourceJsonData.md)
