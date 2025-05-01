---
title: <span class="badge object-type-class"></span> AnnotationTarget
---
# <span class="badge object-type-class"></span> AnnotationTarget

TODO: this should be a regular DataQuery that depends on the selected dashboard

these match the properties of the "grafana" datasouce that is default in most dashboards

## Definition

```java
public class AnnotationTarget {
  public Long limit;
  public Boolean matchAny;
  public List<String> tags;
  public String type;
}
```
## See also

 * <span class="badge builder"></span> [AnnotationTargetBuilder](./builder-AnnotationTargetBuilder.md)
