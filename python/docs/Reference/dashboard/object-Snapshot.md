---
title: <span class="badge object-type-class"></span> Snapshot
---
# <span class="badge object-type-class"></span> Snapshot

A dashboard snapshot shares an interactive dashboard publicly.

It is a read-only version of a dashboard, and is not editable.

It is possible to create a snapshot of a snapshot.

Grafana strips away all sensitive information from the dashboard.

Sensitive information stripped: queries (metric, template,annotation) and panel links.

## Definition

```python
class Snapshot:
    """
    A dashboard snapshot shares an interactive dashboard publicly.
    It is a read-only version of a dashboard, and is not editable.
    It is possible to create a snapshot of a snapshot.
    Grafana strips away all sensitive information from the dashboard.
    Sensitive information stripped: queries (metric, template,annotation) and panel links.
    """

    # Time when the snapshot was created
    created: str
    # Time when the snapshot expires, default is never to expire
    expires: str
    # Is the snapshot saved in an external grafana instance
    external: bool
    # external url, if snapshot was shared in external grafana instance
    external_url: str
    # original url, url of the dashboard that was snapshotted
    original_url: str
    # Unique identifier of the snapshot
    id_val: int
    # Optional, defined the unique key of the snapshot, required if external is true
    key: str
    # Optional, name of the snapshot
    name: str
    # org id of the snapshot
    org_id: int
    # last time when the snapshot was updated
    updated: str
    # url of the snapshot, if snapshot was shared internally
    url: typing.Optional[str]
    # user id of the snapshot creator
    user_id: int
    dashboard: typing.Optional[dashboard.Dashboard]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [Snapshot](./builder-Snapshot.md)
