---
title: <span class="badge object-type-class"></span> LogsQuery
---
# <span class="badge object-type-class"></span> LogsQuery

Azure Monitor Logs sub-query properties

## Definition

```java
public class LogsQuery {
  public String query;
  public ResultFormat resultFormat;
  public List<String> resources;
  public Boolean dashboardTime;
  public String timeColumn;
  public Boolean basicLogsQuery;
  public String workspace;
  public String resource;
  public Boolean intersectTime;
}
```
## See also

 * <span class="badge builder"></span> [LogsQueryBuilder](./builder-LogsQueryBuilder.md)
