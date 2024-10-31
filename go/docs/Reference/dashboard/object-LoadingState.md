---
title: <span class="badge object-type-enum"></span> LoadingState
---
# <span class="badge object-type-enum"></span> LoadingState

Loading status

Accepted values are `NotStarted` (the request is not started), `Loading` (waiting for response), `Streaming` (pulling continuous data), `Done` (response received successfully) or `Error` (failed request).

## Definition

```go
type LoadingState string
const (
	LoadingStateNotStarted LoadingState = "NotStarted"
	LoadingStateLoading LoadingState = "Loading"
	LoadingStateStreaming LoadingState = "Streaming"
	LoadingStateDone LoadingState = "Done"
	LoadingStateError LoadingState = "Error"
)

```
