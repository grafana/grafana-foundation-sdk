---
title: <span class="badge builder"></span> DatasourceVariableBuilder
---
# <span class="badge builder"></span> DatasourceVariableBuilder

## Constructor

```typescript
new DatasourceVariableBuilder(name: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> allFormat

Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.

```typescript
allFormat(allFormat: string)
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

### <span class="badge object-method"></span> id

Unique numeric identifier for the variable.

```typescript
id(id: string)
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

### <span class="badge object-method"></span> regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```typescript
regex(regex: string)
```

### <span class="badge object-method"></span> type

Query used to fetch values for a variable

```typescript
type(query: string | Record<string, any>)
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)
