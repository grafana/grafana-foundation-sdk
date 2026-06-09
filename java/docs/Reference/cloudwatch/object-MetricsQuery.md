---
title: <span class="badge object-type-class"></span> MetricsQuery
---
# <span class="badge object-type-class"></span> MetricsQuery

Shape of a CloudWatch Metrics query

## Definition

```java
public class MetricsQuery {
  public QueryMode queryMode;
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

 * <span class="badge builder"></span> [MetricsQueryBuilder](./builder-MetricsQueryBuilder.md)
