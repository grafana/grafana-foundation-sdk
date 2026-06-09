---
title: <span class="badge object-type-class"></span> LogsQuery
---
# <span class="badge object-type-class"></span> LogsQuery

Shape of a CloudWatch Logs query

## Definition

```java
public class LogsQuery {
  public QueryMode queryMode;
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

 * <span class="badge builder"></span> [LogsQueryBuilder](./builder-LogsQueryBuilder.md)
