---
title: <span class="badge builder"></span> TraceqlFilter
---
# <span class="badge builder"></span> TraceqlFilter

## Constructor

```python
TraceqlFilter()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> tempo.TraceqlFilter
```

### <span class="badge object-method"></span> id_val

Uniquely identify the filter, will not be used in the query generation

```python
def id_val(id_val: str) -> typing.Self
```

### <span class="badge object-method"></span> operator

The operator that connects the tag to the value, for example: =, >, !=, =~

```python
def operator(operator: str) -> typing.Self
```

### <span class="badge object-method"></span> scope

The scope of the filter, can either be unscoped/all scopes, resource or span

```python
def scope(scope: tempo.TraceqlSearchScope) -> typing.Self
```

### <span class="badge object-method"></span> tag

The tag for the search filter, for example: .http.status_code, .service.name, status

```python
def tag(tag: str) -> typing.Self
```

### <span class="badge object-method"></span> value

The value for the search filter

```python
def value(value: typing.Union[str, list[str]]) -> typing.Self
```

### <span class="badge object-method"></span> value_type

The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query

```python
def value_type(value_type: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TraceqlFilter](./object-TraceqlFilter.md)
