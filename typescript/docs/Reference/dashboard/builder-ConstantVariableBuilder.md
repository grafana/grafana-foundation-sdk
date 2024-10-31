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

### <span class="badge object-method"></span> value

Query used to fetch values for a variable

```typescript
value(query: string | Record<string, any>)
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)
