---
title: <span class="badge builder"></span> TimeSeriesList
---
# <span class="badge builder"></span> TimeSeriesList

## Constructor

```python
TimeSeriesList()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> googlecloudmonitoring.TimeSeriesList
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

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```python
def filters(filters: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> group_bys

Array of labels to group data by.

```python
def group_bys(group_bys: list[str]) -> typing.Self
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

### <span class="badge object-method"></span> secondary_alignment_period

Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```python
def secondary_alignment_period(secondary_alignment_period: str) -> typing.Self
```

### <span class="badge object-method"></span> secondary_cross_series_reducer

Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```python
def secondary_cross_series_reducer(secondary_cross_series_reducer: str) -> typing.Self
```

### <span class="badge object-method"></span> secondary_group_bys

Only present if a preprocessor is selected. Array of labels to group data by.

```python
def secondary_group_bys(secondary_group_bys: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> secondary_per_series_aligner

Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.

```python
def secondary_per_series_aligner(secondary_per_series_aligner: str) -> typing.Self
```

### <span class="badge object-method"></span> text

Annotation text.

```python
def text(text: str) -> typing.Self
```

### <span class="badge object-method"></span> title

Annotation title.

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> view

Data view, defaults to FULL.

```python
def view(view: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TimeSeriesList](./object-TimeSeriesList.md)
