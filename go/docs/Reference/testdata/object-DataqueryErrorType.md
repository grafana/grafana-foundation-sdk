---
title: <span class="badge object-type-enum"></span> DataqueryErrorType
---
# <span class="badge object-type-enum"></span> DataqueryErrorType

## Definition

```go
type DataqueryErrorType string
const (
	DataqueryErrorTypeFrontendException DataqueryErrorType = "frontend_exception"
	DataqueryErrorTypeFrontendObservable DataqueryErrorType = "frontend_observable"
	DataqueryErrorTypeServerPanic DataqueryErrorType = "server_panic"
)

```
