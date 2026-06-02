---
title: <span class="badge object-type-class"></span> TracesQuery
---
# <span class="badge object-type-class"></span> TracesQuery

Application Insights Traces sub-query properties

## Definition

```java
public class TracesQuery {
  public ResultFormat resultFormat;
  public List<String> resources;
  public String operationId;
  public List<String> traceTypes;
  public List<TracesFilter> filters;
  public String query;
}
```
## See also

 * <span class="badge builder"></span> [TracesQueryBuilder](./builder-TracesQueryBuilder.md)
