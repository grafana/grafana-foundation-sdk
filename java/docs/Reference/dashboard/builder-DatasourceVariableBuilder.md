---
title: <span class="badge builder"></span> DatasourceVariableBuilder
---
# <span class="badge builder"></span> DatasourceVariableBuilder

## Constructor

```java
new DatasourceVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public DatasourceVariable build()
```

### <span class="badge object-method"></span> allValue

Custom all value

```java
public DatasourceVariableBuilder allValue(String allValue)
```

### <span class="badge object-method"></span> allowCustomValue

Allow custom values to be entered in the variable

```java
public DatasourceVariableBuilder allowCustomValue(Boolean allowCustomValue)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```java
public DatasourceVariableBuilder current(VariableOption current)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```java
public DatasourceVariableBuilder description(String description)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```java
public DatasourceVariableBuilder hide(VariableHide hide)
```

### <span class="badge object-method"></span> includeAll

Whether all value option is available or not

```java
public DatasourceVariableBuilder includeAll(Boolean includeAll)
```

### <span class="badge object-method"></span> label

Optional display name

```java
public DatasourceVariableBuilder label(String label)
```

### <span class="badge object-method"></span> multi

Whether multiple values can be selected or not from variable value list

```java
public DatasourceVariableBuilder multi(Boolean multi)
```

### <span class="badge object-method"></span> name

Name of variable

```java
public DatasourceVariableBuilder name(String name)
```

### <span class="badge object-method"></span> regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```java
public DatasourceVariableBuilder regex(String regex)
```

### <span class="badge object-method"></span> type

Query used to fetch values for a variable

```java
public DatasourceVariableBuilder type(String string)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
