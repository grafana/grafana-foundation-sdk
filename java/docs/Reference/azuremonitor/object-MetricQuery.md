---
title: <span class="badge object-type-class"></span> MetricQuery
---
# <span class="badge object-type-class"></span> MetricQuery

## Definition

```java
public class MetricQuery {
  public List<MonitorResource> resources;
  public String metricNamespace;
  public String customNamespace;
  public String metricName;
  public String region;
  public String timeGrain;
  public String aggregation;
  public List<MetricDimension> dimensionFilters;
  public String top;
  public List<Long> allowedTimeGrainsMs;
  public String alias;
  public String timeGrainUnit;
  public String dimension;
  public String dimensionFilter;
  public String metricDefinition;
  public String resourceUri;
  public String resourceGroup;
  public String resourceName;
}
```
## See also

 * <span class="badge builder"></span> [MetricQueryBuilder](./builder-MetricQueryBuilder.md)
