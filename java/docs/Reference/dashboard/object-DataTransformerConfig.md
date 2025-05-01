---
title: <span class="badge object-type-class"></span> DataTransformerConfig
---
# <span class="badge object-type-class"></span> DataTransformerConfig

Transformations allow to manipulate data returned by a query before the system applies a visualization.

Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,

use the output of one transformation as the input to another transformation, etc.

## Definition

```java
public class DataTransformerConfig {
  public String id;
  public Boolean disabled;
  public MatcherConfig filter;
  public Object options;
}
```
