---
title: <span class="badge builder"></span> ExprTypeReduceSettingsBuilder
---
# <span class="badge builder"></span> ExprTypeReduceSettingsBuilder

## Constructor

```java
new ExprTypeReduceSettingsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ExprTypeReduceSettings build()
```

### <span class="badge object-method"></span> mode

Non-number reduce behavior

Possible enum values:

 - `"dropNN"` Drop non-numbers

 - `"replaceNN"` Replace non-numbers

```java
public ExprTypeReduceSettingsBuilder mode(ExprTypeReduceSettingsMode mode)
```

### <span class="badge object-method"></span> replaceWithValue

Only valid when mode is replace

```java
public ExprTypeReduceSettingsBuilder replaceWithValue(Double replaceWithValue)
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeReduceSettings](./object-ExprTypeReduceSettings.md)
