---
title: <span class="badge object-type-class"></span> CloudMonitoringQuery
---
# <span class="badge object-type-class"></span> CloudMonitoringQuery

## Definition

```java
public class CloudMonitoringQuery extends com.grafana.foundation.cog.variants.Dataquery {
  public String refId;
  public Boolean hide;
  public String queryType;
  public String aliasBy;
  public TimeSeriesList timeSeriesList;
  public TimeSeriesQuery timeSeriesQuery;
  public SLOQuery sloQuery;
  public DataSourceRef datasource;
  public Double intervalMs;
}
```
## See also

 * <span class="badge builder"></span> [CloudMonitoringQueryBuilder](./builder-CloudMonitoringQueryBuilder.md)
