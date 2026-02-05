---
title: <span class="badge builder"></span> NotificationPolicyBuilder
---
# <span class="badge builder"></span> NotificationPolicyBuilder

## Constructor

```java
new NotificationPolicyBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public NotificationPolicy build()
```

### <span class="badge object-method"></span> activeTimeIntervals

```java
public NotificationPolicyBuilder activeTimeIntervals(List<String> activeTimeIntervals)
```

### <span class="badge object-method"></span> continue

```java
public NotificationPolicyBuilder continue(Boolean continueArg)
```

### <span class="badge object-method"></span> groupBy

```java
public NotificationPolicyBuilder groupBy(List<String> groupBy)
```

### <span class="badge object-method"></span> groupInterval

```java
public NotificationPolicyBuilder groupInterval(String groupInterval)
```

### <span class="badge object-method"></span> groupWait

```java
public NotificationPolicyBuilder groupWait(String groupWait)
```

### <span class="badge object-method"></span> match

Deprecated. Remove before v1.0 release.

```java
public NotificationPolicyBuilder match(Map<String, String> match)
```

### <span class="badge object-method"></span> matchRe

```java
public NotificationPolicyBuilder matchRe(Map<String, String> matchRe)
```

### <span class="badge object-method"></span> matchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

```java
public NotificationPolicyBuilder matchers(List<com.grafana.foundation.cog.Builder<Matcher>> matchers)
```

### <span class="badge object-method"></span> muteTimeIntervals

```java
public NotificationPolicyBuilder muteTimeIntervals(List<String> muteTimeIntervals)
```

### <span class="badge object-method"></span> objectMatchers

```java
public NotificationPolicyBuilder objectMatchers(List<List<String>> objectMatchers)
```

### <span class="badge object-method"></span> provenance

```java
public NotificationPolicyBuilder provenance(String provenance)
```

### <span class="badge object-method"></span> receiver

```java
public NotificationPolicyBuilder receiver(String receiver)
```

### <span class="badge object-method"></span> repeatInterval

```java
public NotificationPolicyBuilder repeatInterval(String repeatInterval)
```

### <span class="badge object-method"></span> routes

```java
public NotificationPolicyBuilder routes(List<com.grafana.foundation.cog.Builder<NotificationPolicy>> routes)
```

## See also

 * <span class="badge object-type-class"></span> [NotificationPolicy](./object-NotificationPolicy.md)
