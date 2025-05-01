---
title: <span class="badge builder"></span> CustomVariableBuilder
---
# <span class="badge builder"></span> CustomVariableBuilder

## Constructor

```java
new CustomVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public CustomVariable build()
```

### <span class="badge object-method"></span> allValue

Custom all value

```java
public CustomVariableBuilder allValue(String allValue)
```

### <span class="badge object-method"></span> allowCustomValue

Allow custom values to be entered in the variable

```java
public CustomVariableBuilder allowCustomValue(Boolean allowCustomValue)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```java
public CustomVariableBuilder current(VariableOption current)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```java
public CustomVariableBuilder description(String description)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```java
public CustomVariableBuilder hide(VariableHide hide)
```

### <span class="badge object-method"></span> includeAll

Whether all value option is available or not

```java
public CustomVariableBuilder includeAll(Boolean includeAll)
```

### <span class="badge object-method"></span> label

Optional display name

```java
public CustomVariableBuilder label(String label)
```

### <span class="badge object-method"></span> multi

Whether multiple values can be selected or not from variable value list

```java
public CustomVariableBuilder multi(Boolean multi)
```

### <span class="badge object-method"></span> name

Name of variable

```java
public CustomVariableBuilder name(String name)
```

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```java
public CustomVariableBuilder options(List<VariableOption> options)
```

### <span class="badge object-method"></span> values

Query used to fetch values for a variable

```java
public CustomVariableBuilder values(StringOrMap query)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
