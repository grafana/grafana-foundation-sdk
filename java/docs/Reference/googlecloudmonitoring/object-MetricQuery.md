---
title: <span class="badge object-type-class"></span> MetricQuery
---
# <span class="badge object-type-class"></span> MetricQuery

@deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.

## Definition

```java
public class MetricQuery {
  public String projectName;
  public String perSeriesAligner;
  public String alignmentPeriod;
  public String aliasBy;
  public String editorMode;
  public String metricType;
  public String crossSeriesReducer;
  public List<String> groupBys;
  public List<String> filters;
  public MetricKind metricKind;
  public String valueType;
  public String view;
  public String query;
  public PreprocessorType preprocessor;
  public String graphPeriod;
}
```
## See also

 * <span class="badge builder"></span> [MetricQueryBuilder](./builder-MetricQueryBuilder.md)
