---
title: <span class="badge object-type-class"></span> Panel
---
# <span class="badge object-type-class"></span> Panel

Dashboard panels are the basic visualization building blocks.

## Definition

```java
public class Panel {
  public String type;
  public Integer id;
  public String pluginVersion;
  public List<Dataquery> targets;
  public String title;
  public String description;
  public Boolean transparent;
  public DataSourceRef datasource;
  public GridPos gridPos;
  public List<DashboardLink> links;
  public String repeat;
  public PanelRepeatDirection repeatDirection;
  public Double maxPerRow;
  public Double maxDataPoints;
  public List<DataTransformerConfig> transformations;
  public String interval;
  public String timeFrom;
  public String timeShift;
  public Boolean hideTimeOverride;
  public String timeCompare;
  public LibraryPanelRef libraryPanel;
  public String cacheTimeout;
  public Double queryCachingTTL;
  public Object options;
  public FieldConfigSource fieldConfig;
}
```
## See also

 * <span class="badge builder"></span> [annotationslist.PanelBuilder](../annotationslist/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [barchart.PanelBuilder](../barchart/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [bargauge.PanelBuilder](../bargauge/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [candlestick.PanelBuilder](../candlestick/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [canvas.PanelBuilder](../canvas/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [dashboardlist.PanelBuilder](../dashboardlist/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [datagrid.PanelBuilder](../datagrid/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [debug.PanelBuilder](../debug/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [gauge.PanelBuilder](../gauge/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [geomap.PanelBuilder](../geomap/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [heatmap.PanelBuilder](../heatmap/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [histogram.PanelBuilder](../histogram/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [logs.PanelBuilder](../logs/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [logsnew.PanelBuilder](../logsnew/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [news.PanelBuilder](../news/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [nodegraph.PanelBuilder](../nodegraph/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [piechart.PanelBuilder](../piechart/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [stat.PanelBuilder](../stat/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [statetimeline.PanelBuilder](../statetimeline/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [statushistory.PanelBuilder](../statushistory/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [table.PanelBuilder](../table/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [text.PanelBuilder](../text/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [timeseries.PanelBuilder](../timeseries/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [trend.PanelBuilder](../trend/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [xychart.PanelBuilder](../xychart/builder-PanelBuilder.md)
