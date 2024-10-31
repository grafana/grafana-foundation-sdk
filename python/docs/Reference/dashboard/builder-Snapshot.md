---
title: <span class="badge builder"></span> Snapshot
---
# <span class="badge builder"></span> Snapshot

## Constructor

```python
Snapshot()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.Snapshot
```

### <span class="badge object-method"></span> dashboard

```python
def dashboard(dashboard: cogbuilder.Builder[dashboard.Dashboard]) -> typing.Self
```

### <span class="badge object-method"></span> expires

Time when the snapshot expires, default is never to expire

```python
def expires(expires: str) -> typing.Self
```

### <span class="badge object-method"></span> external

Is the snapshot saved in an external grafana instance

```python
def external(external: bool) -> typing.Self
```

### <span class="badge object-method"></span> external_url

external url, if snapshot was shared in external grafana instance

```python
def external_url(external_url: str) -> typing.Self
```

### <span class="badge object-method"></span> id_val

Unique identifier of the snapshot

```python
def id_val(id_val: int) -> typing.Self
```

### <span class="badge object-method"></span> key

Optional, defined the unique key of the snapshot, required if external is true

```python
def key(key: str) -> typing.Self
```

### <span class="badge object-method"></span> name

Optional, name of the snapshot

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> org_id

org id of the snapshot

```python
def org_id(org_id: int) -> typing.Self
```

### <span class="badge object-method"></span> url

url of the snapshot, if snapshot was shared internally

```python
def url(url: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Snapshot](./object-Snapshot.md)
