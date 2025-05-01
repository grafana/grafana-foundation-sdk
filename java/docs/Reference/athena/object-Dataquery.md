---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts

## Definition

```java
public class Dataquery extends com.grafana.foundation.cog.variants.Dataquery {
  public FormatOptions format;
  public ConnectionArgs connectionArgs;
  public String table;
  public String column;
  public String queryID;
  public String refId;
  public Boolean hide;
  public String queryType;
  public String rawSQL;
  public DataSourceRef datasource;
}
```
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
