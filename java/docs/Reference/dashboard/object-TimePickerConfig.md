---
title: <span class="badge object-type-class"></span> TimePickerConfig
---
# <span class="badge object-type-class"></span> TimePickerConfig

Time picker configuration

It defines the default config for the time picker and the refresh picker for the specific dashboard.

## Definition

```java
public class TimePickerConfig {
  public Boolean hidden;
  public List<String> refreshIntervals;
  public List<TimeOption> quickRanges;
  public String nowDelay;
}
```
## See also

 * <span class="badge builder"></span> [TimePickerBuilder](./builder-TimePickerBuilder.md)
