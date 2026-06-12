---
title: <span class="badge builder"></span> SnapshotBuilder
---
# <span class="badge builder"></span> SnapshotBuilder

A dashboard snapshot shares an interactive dashboard publicly.

It is a read-only version of a dashboard, and is not editable.

It is possible to create a snapshot of a snapshot.

Grafana strips away all sensitive information from the dashboard.

Sensitive information stripped: queries (metric, template,annotation) and panel links.

## Constructor

```typescript
new SnapshotBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> dashboard

```typescript
dashboard(dashboard: cog.Builder<dashboard.Dashboard>)
```

### <span class="badge object-method"></span> expires

Time when the snapshot expires, default is never to expire

```typescript
expires(expires: string)
```

### <span class="badge object-method"></span> external

Is the snapshot saved in an external grafana instance

```typescript
external(external: boolean)
```

### <span class="badge object-method"></span> externalUrl

external url, if snapshot was shared in external grafana instance

```typescript
externalUrl(externalUrl: string)
```

### <span class="badge object-method"></span> id

Unique identifier of the snapshot

```typescript
id(id: number)
```

### <span class="badge object-method"></span> key

Optional, defined the unique key of the snapshot, required if external is true

```typescript
key(key: string)
```

### <span class="badge object-method"></span> name

Optional, name of the snapshot

```typescript
name(name: string)
```

### <span class="badge object-method"></span> orgId

org id of the snapshot

```typescript
orgId(orgId: number)
```

### <span class="badge object-method"></span> originalUrl

original url, url of the dashboard that was snapshotted

```typescript
originalUrl(originalUrl: string)
```

### <span class="badge object-method"></span> url

url of the snapshot, if snapshot was shared internally

```typescript
url(url: string)
```

## See also

 * <span class="badge object-type-interface"></span> [Snapshot](./object-Snapshot.md)
