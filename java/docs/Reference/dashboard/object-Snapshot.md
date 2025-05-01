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

```java
public class Snapshot {
  public String created;
  public String expires;
  public Boolean external;
  public String externalUrl;
  public Integer id;
  public String key;
  public String name;
  public Integer orgId;
  public String updated;
  public String url;
  public Integer userId;
  public Dashboard dashboard;
}
```
## See also

 * <span class="badge builder"></span> [SnapshotBuilder](./builder-SnapshotBuilder.md)
