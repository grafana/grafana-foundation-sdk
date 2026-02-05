---
title: <span class="badge object-type-class"></span> SLOQuery
---
# <span class="badge object-type-class"></span> SLOQuery

SLO sub-query properties.

## Definition

```python
class SLOQuery:
    """
    SLO sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # Alignment function to be used. Defaults to ALIGN_MEAN.
    per_series_aligner: typing.Optional[str]
    # Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignment_period: typing.Optional[str]
    # SLO selector.
    selector_name: str
    # ID for the service the SLO is in.
    service_id: str
    # Name for the service the SLO is in.
    service_name: str
    # ID for the SLO.
    slo_id: str
    # Name of the SLO.
    slo_name: str
    # SLO goal value.
    goal: typing.Optional[float]
    # Specific lookback period for the SLO.
    lookback_period: typing.Optional[str]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [SLOQuery](./builder-SLOQuery.md)
