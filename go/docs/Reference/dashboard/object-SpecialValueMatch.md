---
title: <span class="badge object-type-enum"></span> SpecialValueMatch
---
# <span class="badge object-type-enum"></span> SpecialValueMatch

Special value types supported by the `SpecialValueMap`

## Definition

```go
type SpecialValueMatch string
const (
	SpecialValueMatchTrue SpecialValueMatch = "true"
	SpecialValueMatchFalse SpecialValueMatch = "false"
	SpecialValueMatchNull SpecialValueMatch = "null"
	SpecialValueMatchNaN SpecialValueMatch = "nan"
	SpecialValueMatchNullAndNan SpecialValueMatch = "null+nan"
	SpecialValueMatchEmpty SpecialValueMatch = "empty"
)

```
