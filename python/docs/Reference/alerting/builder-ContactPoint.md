---
title: <span class="badge builder"></span> ContactPoint
---
# <span class="badge builder"></span> ContactPoint

## Constructor

```python
ContactPoint()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alerting.ContactPoint
```

### <span class="badge object-method"></span> disable_resolve_message

```python
def disable_resolve_message(disable_resolve_message: bool) -> typing.Self
```

### <span class="badge object-method"></span> name

Name is used as grouping key in the UI. Contact points with the

same name will be grouped in the UI.

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> provenance

```python
def provenance(provenance: str) -> typing.Self
```

### <span class="badge object-method"></span> settings

```python
def settings(settings: alerting.Json) -> typing.Self
```

### <span class="badge object-method"></span> type_val

```python
def type_val(type_val: typing.Literal["alertmanager", " dingding", " discord", " email", " googlechat", " kafka", " line", " opsgenie", " pagerduty", " pushover", " sensugo", " slack", " teams", " telegram", " threema", " victorops", " webhook", " wecom"]) -> typing.Self
```

### <span class="badge object-method"></span> uid

UID is the unique identifier of the contact point. The UID can be

set by the user.

```python
def uid(uid: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [ContactPoint](./object-ContactPoint.md)
