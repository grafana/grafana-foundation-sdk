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

### <span class="badge object-method"></span> allowCustomValue

Allow custom values to be entered in the variable

```typescript
allowCustomValue(allowCustomValue: boolean)
```

### <span class="badge object-method"></span> datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```typescript
datasource(datasource: common.DataSourceRef)
```

### <span class="badge object-method"></span> definition

```typescript
definition(definition: string)
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

### <span class="badge object-method"></span> staticOptions

Additional static options for query variable

```typescript
staticOptions(staticOptions: dashboard.VariableOption[])
```

### <span class="badge object-method"></span> staticOptionsOrder

Ordering of static options in relation to options returned from data source for query variable

```typescript
staticOptionsOrder(staticOptionsOrder: "before" | "after" | "sorted")
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)
