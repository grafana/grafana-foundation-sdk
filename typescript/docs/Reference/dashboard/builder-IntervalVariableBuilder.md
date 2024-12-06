---
title: <span class="badge builder"></span> IntervalVariableBuilder
---
# <span class="badge builder"></span> IntervalVariableBuilder

## Constructor

```typescript
new IntervalVariableBuilder(name: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> auto

Dynamically calculates interval by dividing time range by the count specified.

```typescript
auto(auto: boolean)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```typescript
current(current: dashboard.VariableOption)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```typescript
description(description: string)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```typescript
hide(hide: dashboard.VariableHide)
```

### <span class="badge object-method"></span> label

Optional display name

```typescript
label(label: string)
```

### <span class="badge object-method"></span> minInterval

The minimum threshold below which the step count intervals will not divide the time.

```typescript
minInterval(autoMin: string)
```

### <span class="badge object-method"></span> name

Name of variable

```typescript
name(name: string)
```

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```typescript
options(options: dashboard.VariableOption[])
```

### <span class="badge object-method"></span> stepCount

How many times the current time range should be divided to calculate the value, similar to the Max data points query option.

For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.

```typescript
stepCount(autoCount: number)
```

### <span class="badge object-method"></span> values

Query used to fetch values for a variable

```typescript
values(query: string | Record<string, any>)
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)
