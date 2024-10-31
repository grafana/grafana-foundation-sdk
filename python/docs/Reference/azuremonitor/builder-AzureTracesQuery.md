---
title: <span class="badge builder"></span> AzureTracesQuery
---
# <span class="badge builder"></span> AzureTracesQuery

## Constructor

```python
AzureTracesQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> azuremonitor.AzureTracesQuery
```

### <span class="badge object-method"></span> filters

Filters for property values.

```python
def filters(filters: list[cogbuilder.Builder[azuremonitor.AzureTracesFilter]]) -> typing.Self
```

### <span class="badge object-method"></span> operation_id

Operation ID. Used only for Traces queries.

```python
def operation_id(operation_id: str) -> typing.Self
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```python
def query(query: str) -> typing.Self
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```python
def resources(resources: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> result_format

Specifies the format results should be returned as.

```python
def result_format(result_format: azuremonitor.ResultFormat) -> typing.Self
```

### <span class="badge object-method"></span> trace_types

Types of events to filter by.

```python
def trace_types(trace_types: list[str]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AzureTracesQuery](./object-AzureTracesQuery.md)
