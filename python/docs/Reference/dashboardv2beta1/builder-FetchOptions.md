---
title: <span class="badge builder"></span> FetchOptions
---
# <span class="badge builder"></span> FetchOptions

## Constructor

```python
FetchOptions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.FetchOptions
```

### <span class="badge object-method"></span> body

```python
def body(body: str) -> typing.Self
```

### <span class="badge object-method"></span> headers

```python
def headers(headers: list[list[str]]) -> typing.Self
```

### <span class="badge object-method"></span> method

```python
def method(method: dashboardv2beta1.HttpRequestMethod) -> typing.Self
```

### <span class="badge object-method"></span> query_params

These are 2D arrays of strings, each representing a key-value pair

We are defining them this way because we can't generate a go struct that

that would have exactly two strings in each sub-array

```python
def query_params(query_params: list[list[str]]) -> typing.Self
```

### <span class="badge object-method"></span> url

```python
def url(url: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [FetchOptions](./object-FetchOptions.md)
