---
title: <span class="badge object-type-enum"></span> DataqueryErrorType
---
# <span class="badge object-type-enum"></span> DataqueryErrorType

## Definition

```go
type DataqueryErrorType string
const (
	DataqueryErrorTypeServerPanic DataqueryErrorType = "server_panic"
	DataqueryErrorTypeFrontendException DataqueryErrorType = "frontend_exception"
	DataqueryErrorTypeFrontendObservable DataqueryErrorType = "frontend_observable"
)

```
