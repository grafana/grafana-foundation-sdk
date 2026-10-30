---
title: <span class="badge builder"></span> AlertRuleInstanceTotals
---
# <span class="badge builder"></span> AlertRuleInstanceTotals

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Constructor

```python
AlertRuleInstanceTotals()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv1beta1.AlertRuleInstanceTotals
```

### <span class="badge object-method"></span> error

```python
def error(error: int) -> typing.Self
```

### <span class="badge object-method"></span> firing

```python
def firing(firing: int) -> typing.Self
```

### <span class="badge object-method"></span> healthy

```python
def healthy(healthy: int) -> typing.Self
```

### <span class="badge object-method"></span> nodata

```python
def nodata(nodata: int) -> typing.Self
```

### <span class="badge object-method"></span> pending

```python
def pending(pending: int) -> typing.Self
```

### <span class="badge object-method"></span> recovering

```python
def recovering(recovering: int) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleInstanceTotals](./object-AlertRuleInstanceTotals.md)
