---
title: <span class="badge object-type-interface"></span> NotificationPolicy
---
# <span class="badge object-type-interface"></span> NotificationPolicy

A Route is a node that contains definitions of how to handle alerts. This is modified

from the upstream alertmanager in that it adds the ObjectMatchers property.

## Definition

```typescript
export interface NotificationPolicy {
	continue?: boolean;
	group_by?: string[];
	group_interval?: string;
	group_wait?: string;
	// Deprecated. Remove before v1.0 release.
	match?: Record<string, string>;
	match_re?: alerting.MatchRegexps;
	// Matchers is a slice of Matchers that is sortable, implements Stringer, and
	// provides a Matches method to match a LabelSet against all Matchers in the
	// slice. Note that some users of Matchers might require it to be sorted.
	matchers?: alerting.Matchers;
	mute_time_intervals?: string[];
	// Matchers is a slice of Matchers that is sortable, implements Stringer, and
	// provides a Matches method to match a LabelSet against all Matchers in the
	// slice. Note that some users of Matchers might require it to be sorted.
	object_matchers?: alerting.ObjectMatchers;
	provenance?: alerting.Provenance;
	receiver?: string;
	repeat_interval?: string;
	routes?: alerting.NotificationPolicy[];
}

```
## See also

 * <span class="badge builder"></span> [NotificationPolicyBuilder](./builder-NotificationPolicyBuilder.md)
