---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```java
public class Dataquery extends com.grafana.foundation.cog.variants.Dataquery {
  public String alias;
  public String query;
  public String timeField;
  public List<BucketAggregation> bucketAggs;
  public List<MetricAggregation> metrics;
  public String refId;
  public Boolean hide;
  public String queryType;
  public DataSourceRef datasource;
}
```
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
