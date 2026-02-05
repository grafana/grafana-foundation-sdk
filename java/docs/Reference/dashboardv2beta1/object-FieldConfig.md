---
title: <span class="badge object-type-class"></span> FieldConfig
---
# <span class="badge object-type-class"></span> FieldConfig

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```java
public class FieldConfig {
  public String displayName;
  public String displayNameFromDS;
  public String description;
  public String path;
  public Boolean writeable;
  public Boolean filterable;
  public String unit;
  public Double decimals;
  public Double min;
  public Double max;
  public List<ValueMapping> mappings;
  public ThresholdsConfig thresholds;
  public FieldColor color;
  public List<Object> links;
  public List<Action> actions;
  public String noValue;
  public Object custom;
}
```
