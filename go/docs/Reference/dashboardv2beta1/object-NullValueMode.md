---
title: <span class="badge object-type-enum"></span> NullValueMode
---
# <span class="badge object-type-enum"></span> NullValueMode

How null values should be handled

## Definition

```go
type NullValueMode string
const (
	NullValueModeNull NullValueMode = "null"
	NullValueModeConnected NullValueMode = "connected"
	NullValueModeNullAsZero NullValueMode = "null as zero"
)

```
