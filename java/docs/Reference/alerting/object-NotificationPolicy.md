---
title: <span class="badge object-type-class"></span> NotificationPolicy
---
# <span class="badge object-type-class"></span> NotificationPolicy

A Route is a node that contains definitions of how to handle alerts. This is modified

from the upstream alertmanager in that it adds the ObjectMatchers property.

## Definition

```java
public class NotificationPolicy {
  public List<String> activeTimeIntervals;
  public Boolean continueArg;
  public List<String> groupBy;
  public String groupInterval;
  public String groupWait;
  public Map<String, String> match;
  public Map<String, String> matchRe;
  public List<Matcher> matchers;
  public List<String> muteTimeIntervals;
  public List<List<String>> objectMatchers;
  public String provenance;
  public String receiver;
  public String repeatInterval;
  public List<NotificationPolicy> routes;
}
```
## See also

 * <span class="badge builder"></span> [NotificationPolicyBuilder](./builder-NotificationPolicyBuilder.md)
