---
title: <span class="badge builder"></span> SLOQuery
---
# <span class="badge builder"></span> SLOQuery

## Constructor

```python
SLOQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> googlecloudmonitoring.SLOQuery
```

### <span class="badge object-method"></span> alignment_period

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```python
def alignment_period(alignment_period: str) -> typing.Self
```

### <span class="badge object-method"></span> goal

SLO goal value.

```python
def goal(goal: float) -> typing.Self
```

### <span class="badge object-method"></span> lookback_period

Specific lookback period for the SLO.

```python
def lookback_period(lookback_period: str) -> typing.Self
```

### <span class="badge object-method"></span> per_series_aligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```python
def per_series_aligner(per_series_aligner: str) -> typing.Self
```

### <span class="badge object-method"></span> project_name

GCP project to execute the query against.

```python
def project_name(project_name: str) -> typing.Self
```

### <span class="badge object-method"></span> selector_name

SLO selector.

```python
def selector_name(selector_name: str) -> typing.Self
```

### <span class="badge object-method"></span> service_id

ID for the service the SLO is in.

```python
def service_id(service_id: str) -> typing.Self
```

### <span class="badge object-method"></span> service_name

Name for the service the SLO is in.

```python
def service_name(service_name: str) -> typing.Self
```

### <span class="badge object-method"></span> slo_id

ID for the SLO.

```python
def slo_id(slo_id: str) -> typing.Self
```

### <span class="badge object-method"></span> slo_name

Name of the SLO.

```python
def slo_name(slo_name: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SLOQuery](./object-SLOQuery.md)
