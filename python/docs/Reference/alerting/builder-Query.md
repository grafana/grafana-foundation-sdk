---
title: <span class="badge builder"></span> Query
---
# <span class="badge builder"></span> Query

## Constructor

```python
Query(ref_id: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alerting.Query
```

### <span class="badge object-method"></span> datasource_uid

Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.

```python
def datasource_uid(datasource_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> model

JSON is the raw JSON query and includes the above properties as well as custom properties.

```python
def model(model: cogbuilder.Builder[cogvariants.Dataquery]) -> typing.Self
```

### <span class="badge object-method"></span> query_type

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

RefID is the unique identifier of the query, set by the frontend call.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> relative_time_range

RelativeTimeRange is the per query start and end time

for requests.

```python
def relative_time_range(relative_time_range: alerting.RelativeTimeRange) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Query](./object-Query.md)
