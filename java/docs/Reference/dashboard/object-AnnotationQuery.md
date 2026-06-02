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
  public Dataquery target;
  public String type;
  public Double builtIn;
  public AnnotationQueryPlacement placement;
  public String expr;
  public String textFormat;
  public String titleFormat;
  public String tagKeys;
  public String step;
  public Boolean useValueForTime;
  public Map<String, AnnotationEventFieldMapping> mappings;
}
```
## See also

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
