---
title: <span class="badge builder"></span> TimeSeriesQuery
---
# <span class="badge builder"></span> TimeSeriesQuery

## Constructor

```python
TimeSeriesQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> googlecloudmonitoring.TimeSeriesQuery
```

### <span class="badge object-method"></span> graph_period

To disable the graphPeriod, it should explictly be set to 'disabled'.

```python
def graph_period(graph_period: str) -> typing.Self
```

### <span class="badge object-method"></span> project_name

GCP project to execute the query against.

```python
def project_name(project_name: str) -> typing.Self
```

### <span class="badge object-method"></span> query

MQL query to be executed.

```python
def query(query: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TimeSeriesQuery](./object-TimeSeriesQuery.md)
