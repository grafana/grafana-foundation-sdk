---
title: <span class="badge object-type-array"></span> Matchers
---
# <span class="badge object-type-array"></span> Matchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

## Definition

```go
type Matchers []alerting.Matcher
```
