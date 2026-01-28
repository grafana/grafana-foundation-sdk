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
def build() -> dashboardv2beta1.Action
```

### <span class="badge object-method"></span> confirmation

```python
def confirmation(confirmation: str) -> typing.Self
```

### <span class="badge object-method"></span> fetch

```python
def fetch(fetch: cogbuilder.Builder[dashboardv2beta1.FetchOptions]) -> typing.Self
```

### <span class="badge object-method"></span> infinity

```python
def infinity(infinity: cogbuilder.Builder[dashboardv2beta1.InfinityOptions]) -> typing.Self
```

### <span class="badge object-method"></span> one_click

```python
def one_click(one_click: bool) -> typing.Self
```

### <span class="badge object-method"></span> style

```python
def style(style: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1ActionStyle]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> type_val

```python
def type_val(type_val: dashboardv2beta1.ActionType) -> typing.Self
```

### <span class="badge object-method"></span> variables

```python
def variables(variables: list[cogbuilder.Builder[dashboardv2beta1.ActionVariable]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Action](./object-Action.md)
