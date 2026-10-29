---
title: <span class="badge object-type-class"></span> RuleSearchHitFields
---
# <span class="badge object-type-class"></span> RuleSearchHitFields

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Definition

```java
public class RuleSearchHitFields {
  public String title;
  public String folder;
  public String type;
  public String interval;
  public Boolean paused;
  public Map<String, String> labels;
  public List<String> datasourceUIDs;
  public Map<String, String> annotations;
  public String forArg;
  public String keepFiringFor;
  public String dashboardUID;
  public Long panelID;
  public String receiver;
  public String notificationType;
  public String routingTree;
  public String metric;
  public String targetDatasourceUID;
}
```
## See also

 * <span class="badge builder"></span> [RuleSearchHitFieldsBuilder](./builder-RuleSearchHitFieldsBuilder.md)
