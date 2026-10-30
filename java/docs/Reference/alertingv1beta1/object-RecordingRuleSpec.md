---
title: <span class="badge object-type-class"></span> RecordingRuleSpec
---
# <span class="badge object-type-class"></span> RecordingRuleSpec

## Definition

```java
public class RecordingRuleSpec {
  public String title;
  public Boolean paused;
  public IntervalTrigger trigger;
  public Map<String, String> labels;
  public String metric;
  public Map<String, Expression> expressions;
  public String targetDatasourceUID;
}
```
## See also

 * <span class="badge builder"></span> [RecordingRuleBuilder](./builder-RecordingRuleBuilder.md)
