---
title: <span class="badge builder"></span> CloudMonitoringQuery
---
# <span class="badge builder"></span> CloudMonitoringQuery

## Constructor

```python
CloudMonitoringQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> googlecloudmonitoring.CloudMonitoringQuery
```

### <span class="badge object-method"></span> alias_by

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```python
def alias_by(alias_by: str) -> typing.Self
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> interval_ms

Time interval in milliseconds.

```python
def interval_ms(interval_ms: float) -> typing.Self
```

### <span class="badge object-method"></span> prom_ql_query

PromQL sub-query properties.

```python
def prom_ql_query(prom_ql_query: cogbuilder.Builder[googlecloudmonitoring.PromQLQuery]) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> slo_query

SLO sub-query properties.

```python
def slo_query(slo_query: cogbuilder.Builder[googlecloudmonitoring.SLOQuery]) -> typing.Self
```

### <span class="badge object-method"></span> time_series_list

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```python
def time_series_list(time_series_list: cogbuilder.Builder[googlecloudmonitoring.TimeSeriesList]) -> typing.Self
```

### <span class="badge object-method"></span> time_series_query

Time Series sub-query properties.

```python
def time_series_query(time_series_query: cogbuilder.Builder[googlecloudmonitoring.TimeSeriesQuery]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CloudMonitoringQuery](./object-CloudMonitoringQuery.md)
