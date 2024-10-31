---
title: <span class="badge builder"></span> NotificationPolicyBuilder
---
# <span class="badge builder"></span> NotificationPolicyBuilder

## Constructor

```typescript
new NotificationPolicyBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> continueVal

```typescript
continueVal(continueVal: boolean)
```

### <span class="badge object-method"></span> groupBy

```typescript
groupBy(groupBy: string[])
```

### <span class="badge object-method"></span> groupInterval

```typescript
groupInterval(groupInterval: string)
```

### <span class="badge object-method"></span> groupWait

```typescript
groupWait(groupWait: string)
```

### <span class="badge object-method"></span> match

Deprecated. Remove before v1.0 release.

```typescript
match(match: Record<string, string>)
```

### <span class="badge object-method"></span> matchRe

```typescript
matchRe(matchRe: alerting.MatchRegexps)
```

### <span class="badge object-method"></span> matchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

```typescript
matchers(matchers: alerting.Matchers)
```

### <span class="badge object-method"></span> muteTimeIntervals

```typescript
muteTimeIntervals(muteTimeIntervals: string[])
```

### <span class="badge object-method"></span> objectMatchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

```typescript
objectMatchers(objectMatchers: alerting.ObjectMatchers)
```

### <span class="badge object-method"></span> provenance

```typescript
provenance(provenance: alerting.Provenance)
```

### <span class="badge object-method"></span> receiver

```typescript
receiver(receiver: string)
```

### <span class="badge object-method"></span> repeatInterval

```typescript
repeatInterval(repeatInterval: string)
```

### <span class="badge object-method"></span> routes

```typescript
routes(routes: cog.Builder<alerting.NotificationPolicy>[])
```

## See also

 * <span class="badge object-type-interface"></span> [NotificationPolicy](./object-NotificationPolicy.md)
