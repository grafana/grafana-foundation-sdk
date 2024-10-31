---
title: <span class="badge object-type-enum"></span> VariableSort
---
# <span class="badge object-type-enum"></span> VariableSort

Sort variable options

Accepted values are:

`0`: No sorting

`1`: Alphabetical ASC

`2`: Alphabetical DESC

`3`: Numerical ASC

`4`: Numerical DESC

`5`: Alphabetical Case Insensitive ASC

`6`: Alphabetical Case Insensitive DESC

`7`: Natural ASC

`8`: Natural DESC

## Definition

```go
type VariableSort int64
const (
	VariableSortDisabled VariableSort = 0
	VariableSortAlphabeticalAsc VariableSort = 1
	VariableSortAlphabeticalDesc VariableSort = 2
	VariableSortNumericalAsc VariableSort = 3
	VariableSortNumericalDesc VariableSort = 4
	VariableSortAlphabeticalCaseInsensitiveAsc VariableSort = 5
	VariableSortAlphabeticalCaseInsensitiveDesc VariableSort = 6
	VariableSortNaturalAsc VariableSort = 7
	VariableSortNaturalDesc VariableSort = 8
)

```
