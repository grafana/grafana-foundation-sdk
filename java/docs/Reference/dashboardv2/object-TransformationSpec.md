---
title: <span class="badge object-type-class"></span> TransformationSpec
---
# <span class="badge object-type-class"></span> TransformationSpec

Transformations allow to manipulate data returned by a query before the system applies a visualization.

Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,

use the output of one transformation as the input to another transformation, etc.

## Definition

```java
public class TransformationSpec {
  public Boolean disabled;
  public MatcherConfig filter;
  public DataTopic topic;
  public Object options;
}
```
