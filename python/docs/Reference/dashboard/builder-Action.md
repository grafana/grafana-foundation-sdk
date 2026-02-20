---
title: <span class="badge builder"></span> Action
---
# <span class="badge builder"></span> Action

## Constructor

```python
Action()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.Action
```

### <span class="badge object-method"></span> confirmation

```python
def confirmation(confirmation: str) -> typing.Self
```

### <span class="badge object-method"></span> fetch

```python
def fetch(fetch: cogbuilder.Builder[dashboard.FetchOptions]) -> typing.Self
```

### <span class="badge object-method"></span> infinity

```python
def infinity(infinity: cogbuilder.Builder[dashboard.InfinityOptions]) -> typing.Self
```

### <span class="badge object-method"></span> one_click

```python
def one_click(one_click: bool) -> typing.Self
```

### <span class="badge object-method"></span> style

```python
def style(style: cogbuilder.Builder[dashboard.DashboardActionStyle]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> type_val

```python
def type_val(type_val: dashboard.ActionType) -> typing.Self
```

### <span class="badge object-method"></span> variables

```python
def variables(variables: list[cogbuilder.Builder[dashboard.ActionVariable]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Action](./object-Action.md)
