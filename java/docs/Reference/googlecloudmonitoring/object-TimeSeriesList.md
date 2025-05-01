---
title: <span class="badge object-type-class"></span> TimeSeriesList
---
# <span class="badge object-type-class"></span> TimeSeriesList

Time Series List sub-query properties.

## Definition

```java
public class TimeSeriesList {
  public String projectName;
  public String crossSeriesReducer;
  public String alignmentPeriod;
  public String perSeriesAligner;
  public List<String> groupBys;
  public List<String> filters;
  public String view;
  public String title;
  public String text;
  public String secondaryCrossSeriesReducer;
  public String secondaryAlignmentPeriod;
  public String secondaryPerSeriesAligner;
  public List<String> secondaryGroupBys;
  public PreprocessorType preprocessor;
}
```
## See also

 * <span class="badge builder"></span> [TimeSeriesListBuilder](./builder-TimeSeriesListBuilder.md)
