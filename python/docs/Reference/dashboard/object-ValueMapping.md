---
title: <span class="badge object-type-disjunction"></span> ValueMapping
---
# <span class="badge object-type-disjunction"></span> ValueMapping

Allow to transform the visual representation of specific data values in a visualization, irrespective of their original units

## Definition

```python
# Allow to transform the visual representation of specific data values in a visualization, irrespective of their original units
ValueMapping: typing.TypeAlias = typing.Union[dashboard.ValueMap, dashboard.RangeMap, dashboard.RegexMap, dashboard.SpecialValueMap]
```
