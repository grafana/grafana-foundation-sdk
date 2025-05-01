---
title: <span class="badge builder"></span> AdHocVariableBuilder
---
# <span class="badge builder"></span> AdHocVariableBuilder

## Constructor

```java
new AdHocVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AdHocVariable build()
```

### <span class="badge object-method"></span> datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```java
public AdHocVariableBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```java
public AdHocVariableBuilder description(String description)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```java
public AdHocVariableBuilder hide(VariableHide hide)
```

### <span class="badge object-method"></span> label

Optional display name

```java
public AdHocVariableBuilder label(String label)
```

### <span class="badge object-method"></span> name

Name of variable

```java
public AdHocVariableBuilder name(String name)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
