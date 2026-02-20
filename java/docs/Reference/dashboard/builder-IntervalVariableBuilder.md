---
title: <span class="badge builder"></span> IntervalVariableBuilder
---
# <span class="badge builder"></span> IntervalVariableBuilder

## Constructor

```java
new IntervalVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public IntervalVariable build()
```

### <span class="badge object-method"></span> allowCustomValue

Allow custom values to be entered in the variable

```java
public IntervalVariableBuilder allowCustomValue(Boolean allowCustomValue)
```

### <span class="badge object-method"></span> auto

Dynamically calculates interval by dividing time range by the count specified.

```java
public IntervalVariableBuilder auto(Boolean auto)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```java
public IntervalVariableBuilder current(VariableOption current)
```

### <span class="badge object-method"></span> definition

```java
public IntervalVariableBuilder definition(String definition)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```java
public IntervalVariableBuilder description(String description)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```java
public IntervalVariableBuilder hide(VariableHide hide)
```

### <span class="badge object-method"></span> label

Optional display name

```java
public IntervalVariableBuilder label(String label)
```

### <span class="badge object-method"></span> minInterval

The minimum threshold below which the step count intervals will not divide the time.

```java
public IntervalVariableBuilder minInterval(String autoMin)
```

### <span class="badge object-method"></span> name

Name of variable

```java
public IntervalVariableBuilder name(String name)
```

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```java
public IntervalVariableBuilder options(List<VariableOption> options)
```

### <span class="badge object-method"></span> staticOptions

Additional static options for query variable

```java
public IntervalVariableBuilder staticOptions(List<VariableOption> staticOptions)
```

### <span class="badge object-method"></span> staticOptionsOrder

Ordering of static options in relation to options returned from data source for query variable

```java
public IntervalVariableBuilder staticOptionsOrder(VariableModelStaticOptionsOrder staticOptionsOrder)
```

### <span class="badge object-method"></span> stepCount

How many times the current time range should be divided to calculate the value, similar to the Max data points query option.

For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.

```java
public IntervalVariableBuilder stepCount(Integer autoCount)
```

### <span class="badge object-method"></span> values

Query used to fetch values for a variable

```java
public IntervalVariableBuilder values(StringOrMap query)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
