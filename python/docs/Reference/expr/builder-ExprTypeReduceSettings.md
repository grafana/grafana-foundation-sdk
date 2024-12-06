---
title: <span class="badge builder"></span> ExprTypeReduceSettings
---
# <span class="badge builder"></span> ExprTypeReduceSettings

## Constructor

```python
ExprTypeReduceSettings()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> expr.ExprTypeReduceSettings
```

### <span class="badge object-method"></span> mode

Non-number reduce behavior

Possible enum values:

 - `"dropNN"` Drop non-numbers

 - `"replaceNN"` Replace non-numbers

```python
def mode(mode: typing.Literal["dropNN", "replaceNN"]) -> typing.Self
```

### <span class="badge object-method"></span> replace_with_value

Only valid when mode is replace

```python
def replace_with_value(replace_with_value: float) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeReduceSettings](./object-ExprTypeReduceSettings.md)
