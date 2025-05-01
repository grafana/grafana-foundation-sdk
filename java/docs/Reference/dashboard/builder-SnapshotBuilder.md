---
title: <span class="badge builder"></span> SnapshotBuilder
---
# <span class="badge builder"></span> SnapshotBuilder

## Constructor

```java
new SnapshotBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Snapshot build()
```

### <span class="badge object-method"></span> dashboard

```java
public SnapshotBuilder dashboard(Dashboard dashboard)
```

### <span class="badge object-method"></span> expires

Time when the snapshot expires, default is never to expire

```java
public SnapshotBuilder expires(String expires)
```

### <span class="badge object-method"></span> external

Is the snapshot saved in an external grafana instance

```java
public SnapshotBuilder external(Boolean external)
```

### <span class="badge object-method"></span> externalUrl

external url, if snapshot was shared in external grafana instance

```java
public SnapshotBuilder externalUrl(String externalUrl)
```

### <span class="badge object-method"></span> id

Unique identifier of the snapshot

```java
public SnapshotBuilder id(Integer id)
```

### <span class="badge object-method"></span> key

Optional, defined the unique key of the snapshot, required if external is true

```java
public SnapshotBuilder key(String key)
```

### <span class="badge object-method"></span> name

Optional, name of the snapshot

```java
public SnapshotBuilder name(String name)
```

### <span class="badge object-method"></span> orgId

org id of the snapshot

```java
public SnapshotBuilder orgId(Integer orgId)
```

### <span class="badge object-method"></span> originalUrl

original url, url of the dashboard that was snapshotted

```java
public SnapshotBuilder originalUrl(String originalUrl)
```

### <span class="badge object-method"></span> url

url of the snapshot, if snapshot was shared internally

```java
public SnapshotBuilder url(String url)
```

## See also

 * <span class="badge object-type-class"></span> [Snapshot](./object-Snapshot.md)
