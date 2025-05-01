---
title: <span class="badge object-type-class"></span> PanelModel
---
# <span class="badge object-type-class"></span> PanelModel

Dashboard panels are the basic visualization building blocks.

## Definition

```java
public class PanelModel {
  public String type;
  public String pluginVersion;
  public List<String> tags;
  public List<Dataquery> targets;
  public String title;
  public String description;
  public Boolean transparent;
  public DataSourceRef datasource;
  public List<DashboardLink> links;
  public String repeat;
  public PanelModelRepeatDirection repeatDirection;
  public Long repeatPanelId;
  public Double maxDataPoints;
  public List<DataTransformerConfig> transformations;
  public String interval;
  public String timeFrom;
  public String timeShift;
  public Object options;
  public FieldConfigSource fieldConfig;
}
```
## See also

 * <span class="badge builder"></span> [PanelModelBuilder](./builder-PanelModelBuilder.md)
