---
title: <span class="badge object-type-enum"></span> DataTopic
---
# <span class="badge object-type-enum"></span> DataTopic

A topic is attached to DataFrame metadata in query results.

This specifies where the data should be used.

## Definition

```go
type DataTopic string
const (
	DataTopicSeries DataTopic = "series"
	DataTopicAnnotations DataTopic = "annotations"
	DataTopicAlertStates DataTopic = "alertStates"
)

```
