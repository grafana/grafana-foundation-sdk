---
title: <span class="badge builder"></span> TextBoxVariableBuilder
---
# <span class="badge builder"></span> TextBoxVariableBuilder

## Constructor

```typescript
new TextBoxVariableBuilder(name: string)
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

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```typescript
current(current: dashboard.VariableOption)
```

### <span class="badge object-method"></span> defaultValue

Query used to fetch values for a variable

```typescript
defaultValue(query: string | Record<string, any>)
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

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```typescript
options(options: dashboard.VariableOption[])
```

## See also

 * <span class="badge object-type-interface"></span> [VariableModel](./object-VariableModel.md)