---
title: <span class="badge builder"></span> QueryVariableBuilder
---
# <span class="badge builder"></span> QueryVariableBuilder

## Constructor

```typescript
new QueryVariableBuilder(name: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> allValue

Custom all value

```typescript
allValue(allValue: string)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```typescript
current(current: dashboard.VariableOption)
```

### <span class="badge object-method"></span> datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```typescript
datasource(datasource: dashboard.DataSourceRef)
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

### <span class="badge object-method"></span> includeAll

Whether all value option is available or not

```typescript
includeAll(includeAll: boolean)
```

### <span class="badge object-method"></span> label

Optional display name

```typescript
label(label: string)
```

### <span class="badge object-method"></span> multi

Whether multiple values can be selected or not from variable value list

```typescript
multi(multi: boolean)
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

### <span class="badge object-method"></span> query

Query used to fetch values for a variable

```typescript
query(query: string | Record<string, any>)
```

### <span class="badge object-method"></span> refresh

Options to config when to refresh a variable

```typescript
refresh(refresh: dashboard.VariableRefresh)
```

### <span class="badge object-method"></span> regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```typescript
regex(regex: string)
```

### <span class="badge object-method"></span> sort

Options sort order

```typescript
sort(sort: dashboard.VariableSort)
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)
