---
title: <span class="badge object-type-enum"></span> TypeSqlType
---
# <span class="badge object-type-enum"></span> TypeSqlType

## Definition

```go
type TypeSqlType string
const (
	TypeSqlTypeNone TypeSqlType = ""
	TypeSqlTypeTimeseriesWide TypeSqlType = "timeseries-wide"
	TypeSqlTypeTimeseriesLong TypeSqlType = "timeseries-long"
	TypeSqlTypeTimeseriesMany TypeSqlType = "timeseries-many"
	TypeSqlTypeTimeseriesMulti TypeSqlType = "timeseries-multi"
	TypeSqlTypeDirectoryListing TypeSqlType = "directory-listing"
	TypeSqlTypeTable TypeSqlType = "table"
	TypeSqlTypeNumericWide TypeSqlType = "numeric-wide"
	TypeSqlTypeNumericMulti TypeSqlType = "numeric-multi"
	TypeSqlTypeNumericLong TypeSqlType = "numeric-long"
	TypeSqlTypeLogLines TypeSqlType = "log-lines"
)

```
