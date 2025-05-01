---
title: <span class="badge object-type-class"></span> ValueMap
---
# <span class="badge object-type-class"></span> ValueMap

Maps text values to a color or different display text and color.

For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.

## Definition

```java
public class ValueMap {
  public MappingType type;
  public Map<String, ValueMappingResult> options;
}
```
