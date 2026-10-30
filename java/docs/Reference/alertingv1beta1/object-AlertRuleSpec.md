---
title: <span class="badge object-type-class"></span> AlertRuleSpec
---
# <span class="badge object-type-class"></span> AlertRuleSpec

## Definition

```java
public class AlertRuleSpec {
  public String title;
  public Boolean paused;
  public IntervalTrigger trigger;
  public Map<String, String> labels;
  public Map<String, String> annotations;
  public String forArg;
  public String keepFiringFor;
  public Long missingSeriesEvalsToResolve;
  public NoDataState noDataState;
  public ExecErrState execErrState;
  public NotificationSettings notificationSettings;
  public Map<String, Expression> expressions;
  public PanelRef panelRef;
}
```
## See also

 * <span class="badge builder"></span> [AlertRuleBuilder](./builder-AlertRuleBuilder.md)
