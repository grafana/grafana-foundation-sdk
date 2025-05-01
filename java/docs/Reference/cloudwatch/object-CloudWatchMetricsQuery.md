---
title: <span class="badge object-type-class"></span> CloudWatchMetricsQuery
---
# <span class="badge object-type-class"></span> CloudWatchMetricsQuery

Shape of a CloudWatch Metrics query

## Definition

```java
public class CloudWatchMetricsQuery extends com.grafana.foundation.cog.variants.Dataquery {
  public CloudWatchQueryMode queryMode;
  public MetricQueryType metricQueryType;
  public MetricEditorMode metricEditorMode;
  public String id;
  public String alias;
  public String label;
  public String expression;
  public String sqlExpression;
  public String refId;
  public Boolean hide;
  public String queryType;
  public String region;
  public String namespace;
  public String metricName;
  public Map<String, StringOrArrayOfString> dimensions;
  public Boolean matchExact;
  public String period;
  public String accountId;
  public String statistic;
  public SQLExpression sql;
  public DataSourceRef datasource;
  public List<String> statistics;
}
```
## See also

 * <span class="badge builder"></span> [CloudWatchMetricsQueryBuilder](./builder-CloudWatchMetricsQueryBuilder.md)
