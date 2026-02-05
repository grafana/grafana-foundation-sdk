---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75

## Definition

```java
public class Dataquery extends com.grafana.foundation.cog.variants.Dataquery {
  public String dataset;
  public String table;
  public String project;
  public QueryFormat format;
  public Boolean rawQuery;
  public String rawSql;
  public String location;
  public Boolean partitioned;
  public String partitionedField;
  public Boolean convertToUTC;
  public Boolean sharded;
  public QueryPriority queryPriority;
  public String timeShift;
  public EditorMode editorMode;
  public SQLExpression sql;
  public String refId;
  public Boolean hide;
  public String queryType;
  public DataSourceRef datasource;
}
```
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
