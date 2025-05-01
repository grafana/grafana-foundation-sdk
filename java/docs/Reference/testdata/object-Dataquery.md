---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```java
public class Dataquery extends com.grafana.foundation.cog.variants.Dataquery {
  public String alias;
  public TestDataQueryType scenarioId;
  public String stringInput;
  public StreamingQuery stream;
  public PulseWaveQuery pulseWave;
  public SimulationQuery sim;
  public List<CSVWave> csvWave;
  public String labels;
  public Long lines;
  public Boolean levelColumn;
  public String channel;
  public NodesQuery nodes;
  public String csvFileName;
  public String csvContent;
  public String rawFrameContent;
  public Integer seriesCount;
  public USAQuery usa;
  public DataqueryErrorType errorType;
  public Integer spanCount;
  public List<List<StringOrInt64>> points;
  public Double dropPercent;
  public String refId;
  public Boolean hide;
  public String queryType;
  public Object datasource;
}
```
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
