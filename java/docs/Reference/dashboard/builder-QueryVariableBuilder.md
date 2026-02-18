---
title: <span class="badge builder"></span> QueryVariableBuilder
---
# <span class="badge builder"></span> QueryVariableBuilder

## Constructor

```java
new QueryVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public QueryVariable build()
```

### <span class="badge object-method"></span> allValue

Custom all value

```java
public QueryVariableBuilder allValue(String allValue)
```

### <span class="badge object-method"></span> allowCustomValue

Allow custom values to be entered in the variable

```java
public QueryVariableBuilder allowCustomValue(Boolean allowCustomValue)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```java
public QueryVariableBuilder current(VariableOption current)
```

### <span class="badge object-method"></span> datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```java
public QueryVariableBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> definition

```java
public QueryVariableBuilder definition(String definition)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```java
public QueryVariableBuilder description(String description)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```java
public QueryVariableBuilder hide(VariableHide hide)
```

### <span class="badge object-method"></span> includeAll

Whether all value option is available or not

```java
public QueryVariableBuilder includeAll(Boolean includeAll)
```

### <span class="badge object-method"></span> label

Optional display name

```java
public QueryVariableBuilder label(String label)
```

### <span class="badge object-method"></span> multi

Whether multiple values can be selected or not from variable value list

```java
public QueryVariableBuilder multi(Boolean multi)
```

### <span class="badge object-method"></span> name

Name of variable

```java
public QueryVariableBuilder name(String name)
```

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```java
public QueryVariableBuilder options(List<VariableOption> options)
```

### <span class="badge object-method"></span> query

Query used to fetch values for a variable

```java
public QueryVariableBuilder query(StringOrMap query)
```

### <span class="badge object-method"></span> refresh

Options to config when to refresh a variable

```java
public QueryVariableBuilder refresh(VariableRefresh refresh)
```

### <span class="badge object-method"></span> regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```java
public QueryVariableBuilder regex(String regex)
```

### <span class="badge object-method"></span> sort

Options sort order

```java
public QueryVariableBuilder sort(VariableSort sort)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
