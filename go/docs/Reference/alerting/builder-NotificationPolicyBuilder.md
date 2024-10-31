---
title: <span class="badge builder"></span> NotificationPolicyBuilder
---
# <span class="badge builder"></span> NotificationPolicyBuilder

## Constructor

```go
func NewNotificationPolicyBuilder() *NotificationPolicyBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NotificationPolicyBuilder) Build() (NotificationPolicy, error)
```

### <span class="badge object-method"></span> Continue

```go
func (builder *NotificationPolicyBuilder) Continue(continueArg bool) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> GroupBy

```go
func (builder *NotificationPolicyBuilder) GroupBy(groupBy []string) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> GroupInterval

```go
func (builder *NotificationPolicyBuilder) GroupInterval(groupInterval string) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> GroupWait

```go
func (builder *NotificationPolicyBuilder) GroupWait(groupWait string) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> Match

Deprecated. Remove before v1.0 release.

```go
func (builder *NotificationPolicyBuilder) Match(match map[string]string) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> MatchRe

```go
func (builder *NotificationPolicyBuilder) MatchRe(matchRe alerting.MatchRegexps) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> Matchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

```go
func (builder *NotificationPolicyBuilder) Matchers(matchers alerting.Matchers) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> MuteTimeIntervals

```go
func (builder *NotificationPolicyBuilder) MuteTimeIntervals(muteTimeIntervals []string) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> ObjectMatchers

```go
func (builder *NotificationPolicyBuilder) ObjectMatchers(objectMatchers alerting.ObjectMatchers) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> Provenance

```go
func (builder *NotificationPolicyBuilder) Provenance(provenance alerting.Provenance) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> Receiver

```go
func (builder *NotificationPolicyBuilder) Receiver(receiver string) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> RepeatInterval

```go
func (builder *NotificationPolicyBuilder) RepeatInterval(repeatInterval string) *NotificationPolicyBuilder
```

### <span class="badge object-method"></span> Routes

```go
func (builder *NotificationPolicyBuilder) Routes(routes []cog.Builder[alerting.NotificationPolicy]) *NotificationPolicyBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [NotificationPolicy](./object-NotificationPolicy.md)
