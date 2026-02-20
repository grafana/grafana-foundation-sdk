---
title: <span class="badge object-type-class"></span> VariableModel
---
# <span class="badge object-type-class"></span> VariableModel

A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.

## Definition

```java
public class VariableModel {
  public VariableType type;
  public String name;
  public String label;
  public VariableHide hide;
  public Boolean skipUrlSync;
  public String description;
  public StringOrMap query;
  public DataSourceRef datasource;
  public VariableOption current;
  public Boolean multi;
  public Boolean allowCustomValue;
  public List<VariableOption> options;
  public VariableRefresh refresh;
  public VariableSort sort;
  public Boolean includeAll;
  public String allValue;
  public String regex;
  public List<VariableOption> staticOptions;
  public VariableModelStaticOptionsOrder staticOptionsOrder;
  public Boolean auto;
  public String autoMin;
  public Integer autoCount;
  public String definition;
}
```
## See also

 * <span class="badge builder"></span> [AdHocVariableBuilder](./builder-AdHocVariableBuilder.md)
 * <span class="badge builder"></span> [ConstantVariableBuilder](./builder-ConstantVariableBuilder.md)
 * <span class="badge builder"></span> [CustomVariableBuilder](./builder-CustomVariableBuilder.md)
 * <span class="badge builder"></span> [DatasourceVariableBuilder](./builder-DatasourceVariableBuilder.md)
 * <span class="badge builder"></span> [IntervalVariableBuilder](./builder-IntervalVariableBuilder.md)
 * <span class="badge builder"></span> [QueryVariableBuilder](./builder-QueryVariableBuilder.md)
 * <span class="badge builder"></span> [TextBoxVariableBuilder](./builder-TextBoxVariableBuilder.md)
