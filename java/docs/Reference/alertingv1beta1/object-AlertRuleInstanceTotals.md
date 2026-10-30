---
title: <span class="badge object-type-class"></span> AlertRuleInstanceTotals
---
# <span class="badge object-type-class"></span> AlertRuleInstanceTotals

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Definition

```java
public class AlertRuleInstanceTotals {
  public Long healthy;
  public Long firing;
  public Long pending;
  public Long recovering;
  public Long nodata;
  public Long error;
}
```
## See also

 * <span class="badge builder"></span> [AlertRuleInstanceTotalsBuilder](./builder-AlertRuleInstanceTotalsBuilder.md)
