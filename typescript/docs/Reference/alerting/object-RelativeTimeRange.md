---
title: <span class="badge object-type-interface"></span> RelativeTimeRange
---
# <span class="badge object-type-interface"></span> RelativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

## Definition

```typescript
export interface RelativeTimeRange {
	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	from?: alerting.Duration;
	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	to?: alerting.Duration;
}

```
## Methods

No methods.
