---
title: <span class="badge builder"></span> PromQLQuery
---
# <span class="badge builder"></span> PromQLQuery

## Constructor

```python
PromQLQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> googlecloudmonitoring.PromQLQuery
```

### <span class="badge object-method"></span> expr

PromQL expression/query to be executed.

```python
def expr(expr: str) -> typing.Self
```

### <span class="badge object-method"></span> project_name

GCP project to execute the query against.

```python
def project_name(project_name: str) -> typing.Self
```

### <span class="badge object-method"></span> step

PromQL min step

```python
def step(step: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [PromQLQuery](./object-PromQLQuery.md)
