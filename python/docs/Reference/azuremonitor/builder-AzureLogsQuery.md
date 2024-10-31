---
title: <span class="badge builder"></span> AzureLogsQuery
---
# <span class="badge builder"></span> AzureLogsQuery

## Constructor

```python
AzureLogsQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> azuremonitor.AzureLogsQuery
```

### <span class="badge object-method"></span> intersect_time

If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false

```python
def intersect_time(intersect_time: bool) -> typing.Self
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```python
def query(query: str) -> typing.Self
```

### <span class="badge object-method"></span> resource

@deprecated Use resources instead

```python
def resource(resource: str) -> typing.Self
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

### <span class="badge object-method"></span> workspace

Workspace ID. This was removed in Grafana 8, but remains for backwards compat

```python
def workspace(workspace: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AzureLogsQuery](./object-AzureLogsQuery.md)
