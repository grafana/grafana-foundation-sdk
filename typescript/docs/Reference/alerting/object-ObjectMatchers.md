---
title: <span class="badge object-type-ref"></span> ObjectMatchers
---
# <span class="badge object-type-ref"></span> ObjectMatchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

## Definition

```typescript
export type ObjectMatchers = alerting.Matchers;

```
