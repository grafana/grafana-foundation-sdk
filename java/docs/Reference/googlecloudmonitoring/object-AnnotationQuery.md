---
title: <span class="badge object-type-class"></span> AnnotationQuery
---
# <span class="badge object-type-class"></span> AnnotationQuery

Annotation sub-query properties.

## Definition

```java
public class AnnotationQuery {
  public String projectName;
  public String crossSeriesReducer;
  public String alignmentPeriod;
  public String perSeriesAligner;
  public List<String> groupBys;
  public List<String> filters;
  public String view;
  public String secondaryCrossSeriesReducer;
  public String secondaryAlignmentPeriod;
  public String secondaryPerSeriesAligner;
  public List<String> secondaryGroupBys;
  public String title;
  public PreprocessorType preprocessor;
  public String text;
}
```
## See also

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
