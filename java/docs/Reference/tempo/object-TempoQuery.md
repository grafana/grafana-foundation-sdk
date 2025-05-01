---
title: <span class="badge object-type-class"></span> TempoQuery
---
# <span class="badge object-type-class"></span> TempoQuery

## Definition

```java
public class TempoQuery extends com.grafana.foundation.cog.variants.Dataquery {
  public String refId;
  public Boolean hide;
  public String queryType;
  public String query;
  public String search;
  public String serviceName;
  public String spanName;
  public String minDuration;
  public String maxDuration;
  public StringOrArrayOfString serviceMapQuery;
  public Boolean serviceMapIncludeNamespace;
  public Long limit;
  public Long spss;
  public List<TraceqlFilter> filters;
  public List<TraceqlFilter> groupBy;
  public SearchTableType tableType;
  public String step;
  public DataSourceRef datasource;
  public Long exemplars;
}
```
## See also

 * <span class="badge builder"></span> [TempoQueryBuilder](./builder-TempoQueryBuilder.md)
