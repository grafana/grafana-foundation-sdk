---
title: <span class="badge object-type-class"></span> AzureLogsQuery
---
# <span class="badge object-type-class"></span> AzureLogsQuery

Azure Monitor Logs sub-query properties

## Definition

```java
public class AzureLogsQuery {
  public String query;
  public ResultFormat resultFormat;
  public List<String> resources;
  public Boolean dashboardTime;
  public String timeColumn;
  public String workspace;
  public String resource;
  public Boolean intersectTime;
}
```
## See also

 * <span class="badge builder"></span> [AzureLogsQueryBuilder](./builder-AzureLogsQueryBuilder.md)
