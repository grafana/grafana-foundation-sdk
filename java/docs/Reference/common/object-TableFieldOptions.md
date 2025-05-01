---
title: <span class="badge object-type-class"></span> TableFieldOptions
---
# <span class="badge object-type-class"></span> TableFieldOptions

Field options for each field within a table (e.g 10, "The String", 64.20, etc.)

Generally defines alignment, filtering capabilties, display options, etc.

## Definition

```java
public class TableFieldOptions {
  public Double width;
  public Double minWidth;
  public FieldTextAlignment align;
  public TableCellDisplayMode displayMode;
  public TableCellOptions cellOptions;
  public Boolean hidden;
  public Boolean inspect;
  public Boolean filterable;
  public Boolean hideHeader;
}
```
## See also

 * <span class="badge builder"></span> [TableFieldOptionsBuilder](./builder-TableFieldOptionsBuilder.md)
