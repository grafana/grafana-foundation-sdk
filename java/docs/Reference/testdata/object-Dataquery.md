---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```java
public class Dataquery extends com.grafana.foundation.cog.variants.Dataquery {
  public String alias;
  public String channel;
  public String csvContent;
  public String csvFileName;
  public List<CSVWave> csvWave;
  public DataSourceRef datasource;
  public Double dropPercent;
  public DataqueryErrorSource errorSource;
  public DataqueryErrorType errorType;
  public Boolean flamegraphDiff;
  public Boolean hide;
  public Double intervalMs;
  public String labels;
  public Boolean levelColumn;
  public Long lines;
  public Double max;
  public Long maxDataPoints;
  public Double min;
  public NodesQuery nodes;
  public Double noise;
  public List<List<Object>> points;
  public PulseWaveQuery pulseWave;
  public String queryType;
  public String rawFrameContent;
  public String refId;
  public ResultAssertions resultAssertions;
  public DataqueryScenarioId scenarioId;
  public Long seriesCount;
  public SimulationQuery sim;
  public Long spanCount;
  public Double spread;
  public Double startValue;
  public StreamingQuery stream;
  public String stringInput;
  public TimeRange timeRange;
  public USAQuery usa;
  public Boolean withNil;
}
```
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
