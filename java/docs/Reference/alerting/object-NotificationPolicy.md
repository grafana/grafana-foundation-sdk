---
title: <span class="badge object-type-class"></span> NotificationPolicy
---
# <span class="badge object-type-class"></span> NotificationPolicy

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
  public String receiver;
  public String repeatInterval;
  public List<NotificationPolicy> routes;
}
```
## See also

 * <span class="badge builder"></span> [NotificationPolicyBuilder](./builder-NotificationPolicyBuilder.md)
