---
title: <span class="badge object-type-class"></span> CloudWatchLogsQuery
---
# <span class="badge object-type-class"></span> CloudWatchLogsQuery

Shape of a CloudWatch Logs query

## Definition

```java
public class CloudWatchLogsQuery extends com.grafana.foundation.cog.variants.Dataquery {
  public CloudWatchQueryMode queryMode;
  public String id;
  public String region;
  public String expression;
  public List<String> statsGroups;
  public List<LogGroup> logGroups;
  public List<String> logGroupNames;
  public String refId;
  public Boolean hide;
  public String queryType;
  public LogsQueryLanguage queryLanguage;
  public DataSourceRef datasource;
}
```
## See also

 * <span class="badge builder"></span> [CloudWatchLogsQueryBuilder](./builder-CloudWatchLogsQueryBuilder.md)
