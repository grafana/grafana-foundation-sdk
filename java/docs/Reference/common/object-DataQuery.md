---
title: <span class="badge object-type-class"></span> DataQuery
---
# <span class="badge object-type-class"></span> DataQuery

These are the common properties available to all queries in all datasources.

Specific implementations will *extend* this interface, adding the required

properties for the given context.

## Definition

```java
public class DataQuery {
  public String refId;
  public Boolean hide;
  public String queryType;
  public Object datasource;
}
```
## See also

 * <span class="badge builder"></span> [DataQueryBuilder](./builder-DataQueryBuilder.md)
