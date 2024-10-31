---
title: <span class="badge builder"></span> MetricQuery
---
# <span class="badge builder"></span> MetricQuery

## Constructor

```python
MetricQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> googlecloudmonitoring.MetricQuery
```

### <span class="badge object-method"></span> alias_by

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```python
def alias_by(alias_by: str) -> typing.Self
```

### <span class="badge object-method"></span> alignment_period

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```python
def alignment_period(alignment_period: str) -> typing.Self
```

### <span class="badge object-method"></span> cross_series_reducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```python
def cross_series_reducer(cross_series_reducer: str) -> typing.Self
```

### <span class="badge object-method"></span> editor_mode

```python
def editor_mode(editor_mode: str) -> typing.Self
```

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```python
def filters(filters: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> graph_period

To disable the graphPeriod, it should explictly be set to 'disabled'.

```python
def graph_period(graph_period: str) -> typing.Self
```

### <span class="badge object-method"></span> group_bys

Array of labels to group data by.

```python
def group_bys(group_bys: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> metric_kind

```python
def metric_kind(metric_kind: googlecloudmonitoring.MetricKind) -> typing.Self
```

### <span class="badge object-method"></span> metric_type

```python
def metric_type(metric_type: str) -> typing.Self
```

### <span class="badge object-method"></span> per_series_aligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```python
def per_series_aligner(per_series_aligner: str) -> typing.Self
```

### <span class="badge object-method"></span> preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```python
def preprocessor(preprocessor: googlecloudmonitoring.PreprocessorType) -> typing.Self
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

### <span class="badge object-method"></span> value_type

```python
def value_type(value_type: str) -> typing.Self
```

### <span class="badge object-method"></span> view

```python
def view(view: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [MetricQuery](./object-MetricQuery.md)
