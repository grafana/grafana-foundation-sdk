---
title: <span class="badge object-type-class"></span> CloudWatchAnnotationQuery
---
# <span class="badge object-type-class"></span> CloudWatchAnnotationQuery

Shape of a CloudWatch Annotation query

TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer

#CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")

## Definition

```java
public class CloudWatchAnnotationQuery extends com.grafana.foundation.cog.variants.Dataquery {
  public CloudWatchQueryMode queryMode;
  public Boolean prefixMatching;
  public String actionPrefix;
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
  public String alarmNamePrefix;
  public DataSourceRef datasource;
  public List<String> statistics;
}
```
## See also

 * <span class="badge builder"></span> [CloudWatchAnnotationQueryBuilder](./builder-CloudWatchAnnotationQueryBuilder.md)
