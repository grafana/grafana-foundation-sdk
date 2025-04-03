---
title: <span class="badge builder"></span> NotificationPolicy
---
# <span class="badge builder"></span> NotificationPolicy

## Constructor

```python
NotificationPolicy()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alerting.NotificationPolicy
```

### <span class="badge object-method"></span> active_time_intervals

```python
def active_time_intervals(active_time_intervals: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> continue_val

```python
def continue_val(continue_val: bool) -> typing.Self
```

### <span class="badge object-method"></span> group_by

```python
def group_by(group_by: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> group_interval

```python
def group_interval(group_interval: str) -> typing.Self
```

### <span class="badge object-method"></span> group_wait

```python
def group_wait(group_wait: str) -> typing.Self
```

### <span class="badge object-method"></span> match

Deprecated. Remove before v1.0 release.

```python
def match(match: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> match_re

```python
def match_re(match_re: alerting.MatchRegexps) -> typing.Self
```

### <span class="badge object-method"></span> matchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

```python
def matchers(matchers: alerting.Matchers) -> typing.Self
```

### <span class="badge object-method"></span> mute_time_intervals

```python
def mute_time_intervals(mute_time_intervals: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> object_matchers

```python
def object_matchers(object_matchers: alerting.ObjectMatchers) -> typing.Self
```

### <span class="badge object-method"></span> provenance

```python
def provenance(provenance: alerting.Provenance) -> typing.Self
```

### <span class="badge object-method"></span> receiver

```python
def receiver(receiver: str) -> typing.Self
```

### <span class="badge object-method"></span> repeat_interval

```python
def repeat_interval(repeat_interval: str) -> typing.Self
```

### <span class="badge object-method"></span> routes

```python
def routes(routes: list[cogbuilder.Builder[alerting.NotificationPolicy]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [NotificationPolicy](./object-NotificationPolicy.md)
