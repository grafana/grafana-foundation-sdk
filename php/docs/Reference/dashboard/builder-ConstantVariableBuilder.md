---
title: <span class="badge builder"></span> ConstantVariableBuilder
---
# <span class="badge builder"></span> ConstantVariableBuilder

## Constructor

```php
new ConstantVariableBuilder(string $name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> allFormat

Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.

```php
allFormat(string $allFormat)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```php
description(string $description)
```

### <span class="badge object-method"></span> id

Unique numeric identifier for the variable.

```php
id(string $id)
```

### <span class="badge object-method"></span> label

Optional display name

```php
label(string $label)
```

### <span class="badge object-method"></span> name

Name of variable

```php
name(string $name)
```

### <span class="badge object-method"></span> value

Query used to fetch values for a variable

@param string|array<string, mixed> $query

```php
value($query)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
