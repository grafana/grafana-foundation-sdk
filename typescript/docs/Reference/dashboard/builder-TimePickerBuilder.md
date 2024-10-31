---
title: <span class="badge builder"></span> TimePickerBuilder
---
# <span class="badge builder"></span> TimePickerBuilder

## Constructor

```typescript
new TimePickerBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> hidden

Whether timepicker is visible or not.

```typescript
hidden(hidden: boolean)
```

### <span class="badge object-method"></span> nowDelay

Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.

```typescript
nowDelay(nowDelay: string)
```

### <span class="badge object-method"></span> refreshIntervals

Interval options available in the refresh picker dropdown.

```typescript
refreshIntervals(refreshIntervals: string[])
```

### <span class="badge object-method"></span> timeOptions

Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.

```typescript
timeOptions(timeOptions: string[])
```

## See also

 * <span class="badge object-type-interface"></span> [TimePickerConfig](./object-TimePickerConfig.md)
