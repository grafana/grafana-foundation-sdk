---
title: <span class="badge builder"></span> ConstantVariableBuilder
---
# <span class="badge builder"></span> ConstantVariableBuilder

## Constructor

```java
new ConstantVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ConstantVariable build()
```

### <span class="badge object-method"></span> allowCustomValue

Allow custom values to be entered in the variable

```java
public ConstantVariableBuilder allowCustomValue(Boolean allowCustomValue)
```

### <span class="badge object-method"></span> definition

```java
public ConstantVariableBuilder definition(String definition)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```java
public ConstantVariableBuilder description(String description)
```

### <span class="badge object-method"></span> label

Optional display name

```java
public ConstantVariableBuilder label(String label)
```

### <span class="badge object-method"></span> name

Name of variable

```java
public ConstantVariableBuilder name(String name)
```

### <span class="badge object-method"></span> staticOptions

Additional static options for query variable

```java
public ConstantVariableBuilder staticOptions(List<VariableOption> staticOptions)
```

### <span class="badge object-method"></span> staticOptionsOrder

Ordering of static options in relation to options returned from data source for query variable

```java
public ConstantVariableBuilder staticOptionsOrder(VariableModelStaticOptionsOrder staticOptionsOrder)
```

### <span class="badge object-method"></span> value

Query used to fetch values for a variable

```java
public ConstantVariableBuilder value(StringOrMap query)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
