---
title: <span class="badge object-type-class"></span> FieldConfigSource
---
# <span class="badge object-type-class"></span> FieldConfigSource

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```java
public class FieldConfigSource {
  public FieldConfig defaults;
  public List<Dashboardv2beta1FieldConfigSourceOverrides> overrides;
}
```
