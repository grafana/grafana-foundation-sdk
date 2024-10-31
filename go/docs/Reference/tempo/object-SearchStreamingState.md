---
title: <span class="badge object-type-enum"></span> SearchStreamingState
---
# <span class="badge object-type-enum"></span> SearchStreamingState

The state of the TraceQL streaming search query

## Definition

```go
type SearchStreamingState string
const (
	SearchStreamingStatePending SearchStreamingState = "pending"
	SearchStreamingStateStreaming SearchStreamingState = "streaming"
	SearchStreamingStateDone SearchStreamingState = "done"
	SearchStreamingStateError SearchStreamingState = "error"
)

```
