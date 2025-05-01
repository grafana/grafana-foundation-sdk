---
title: <span class="badge object-type-class"></span> SpecialValueMap
---
# <span class="badge object-type-class"></span> SpecialValueMap

Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.

See SpecialValueMatch to see the list of special values.

For example, you can configure a special value mapping so that null values appear as N/A.

## Definition

```java
public class SpecialValueMap {
  public MappingType type;
  public DashboardSpecialValueMapOptions options;
}
```
