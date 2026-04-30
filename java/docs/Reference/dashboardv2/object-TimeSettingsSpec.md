---
title: <span class="badge object-type-class"></span> TimeSettingsSpec
---
# <span class="badge object-type-class"></span> TimeSettingsSpec

Time configuration

It defines the default time config for the time picker, the refresh picker for the specific dashboard.

## Definition

```java
public class TimeSettingsSpec {
  public String timezone;
  public String from;
  public String to;
  public String autoRefresh;
  public List<String> autoRefreshIntervals;
  public List<TimeRangeOption> quickRanges;
  public Boolean hideTimepicker;
  public TimeSettingsSpecWeekStart weekStart;
  public Long fiscalYearStartMonth;
  public String nowDelay;
}
```
## See also

 * <span class="badge builder"></span> [TimeSettingsBuilder](./builder-TimeSettingsBuilder.md)
