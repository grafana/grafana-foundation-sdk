---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```java
public class Dataquery extends com.grafana.foundation.cog.variants.Dataquery {
  public List<AdhocFilters> adhocFilters;
  public DataSourceRef datasource;
  public QueryEditorMode editorMode;
  public Boolean exemplar;
  public String expr;
  public PromQueryFormat format;
  public List<String> groupByKeys;
  public Boolean hide;
  public Boolean instant;
  public Long intervalFactor;
  public Double intervalMs;
  public String legendFormat;
  public Long maxDataPoints;
  public String queryType;
  public Boolean range;
  public String refId;
  public ResultAssertions resultAssertions;
  public List<Scopes> scopes;
  public TimeRange timeRange;
  public String interval;
}
```
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
