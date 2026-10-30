---
title: <span class="badge object-type-enum"></span> RecordingRuleHealth
---
# <span class="badge object-type-enum"></span> RecordingRuleHealth

## Definition

```go
type RecordingRuleHealth string
const (
	RecordingRuleHealthUnknown RecordingRuleHealth = "Unknown"
	RecordingRuleHealthRecording RecordingRuleHealth = "Recording"
	RecordingRuleHealthPaused RecordingRuleHealth = "Paused"
	RecordingRuleHealthError RecordingRuleHealth = "Error"
	RecordingRuleHealthNoData RecordingRuleHealth = "NoData"
)

```
