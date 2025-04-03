---
title: <span class="badge object-type-class"></span> RelativeTimeRange
---
# <span class="badge object-type-class"></span> RelativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

## Definition

```python
class RelativeTimeRange:
    """
    RelativeTimeRange is the per query start and end time
    for requests.
    """

    # A Duration represents the elapsed time between two instants
    # as an int64 nanosecond count. The representation limits the
    # largest representable duration to approximately 290 years.
    from_val: typing.Optional[alerting.Duration]
    # A Duration represents the elapsed time between two instants
    # as an int64 nanosecond count. The representation limits the
    # largest representable duration to approximately 290 years.
    to: typing.Optional[alerting.Duration]
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

