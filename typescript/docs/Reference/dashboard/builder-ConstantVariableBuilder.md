---
title: <span class="badge builder"></span> ConstantVariableBuilder
---
# <span class="badge builder"></span> ConstantVariableBuilder

## Constructor

```typescript
new ConstantVariableBuilder(name: string)
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

### <span class="badge object-method"></span> definition

```typescript
definition(definition: string)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```typescript
description(description: string)
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

### <span class="badge object-method"></span> value

Query used to fetch values for a variable

```typescript
value(query: string | Record<string, any>)
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)
