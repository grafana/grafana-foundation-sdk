---
title: <span class="badge object-type-class"></span> AnnotationQuery
---
# <span class="badge object-type-class"></span> AnnotationQuery

TODO docs

FROM: AnnotationQuery in grafana-data/src/types/annotations.ts

## Definition

```java
public class AnnotationQuery {
  public String name;
  public DataSourceRef datasource;
  public Boolean enable;
  public Boolean hide;
  public String iconColor;
  public AnnotationPanelFilter filter;
  public AnnotationTarget target;
  public String type;
  public Double builtIn;
  public String expr;
}
```
## See also

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
