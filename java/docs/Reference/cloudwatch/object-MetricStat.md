---
title: <span class="badge object-type-class"></span> MetricStat
---
# <span class="badge object-type-class"></span> MetricStat

## Definition

```java
public class MetricStat {
  public String region;
  public String namespace;
  public String metricName;
  public Map<String, StringOrArrayOfString> dimensions;
  public Boolean matchExact;
  public String period;
  public String accountId;
  public String statistic;
  public List<String> statistics;
}
```
## See also

 * <span class="badge builder"></span> [MetricStatBuilder](./builder-MetricStatBuilder.md)
