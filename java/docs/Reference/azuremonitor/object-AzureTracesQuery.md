---
title: <span class="badge object-type-class"></span> AzureTracesQuery
---
# <span class="badge object-type-class"></span> AzureTracesQuery

Application Insights Traces sub-query properties

## Definition

```java
public class AzureTracesQuery {
  public ResultFormat resultFormat;
  public List<String> resources;
  public String operationId;
  public List<String> traceTypes;
  public List<AzureTracesFilter> filters;
  public String query;
}
```
## See also

 * <span class="badge builder"></span> [AzureTracesQueryBuilder](./builder-AzureTracesQueryBuilder.md)
