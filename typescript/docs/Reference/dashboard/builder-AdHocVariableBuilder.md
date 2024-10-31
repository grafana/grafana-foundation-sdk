---
title: <span class="badge builder"></span> AdHocVariableBuilder
---
# <span class="badge builder"></span> AdHocVariableBuilder

## Constructor

```typescript
new AdHocVariableBuilder(name: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
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

### <span class="badge object-method"></span> label

Optional display name

```typescript
label(label: string)
```

### <span class="badge object-method"></span> name

Name of variable

```typescript
name(name: string)
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)
