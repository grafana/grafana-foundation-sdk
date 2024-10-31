---
title: <span class="badge object-type-map"></span> Dimensions
---
# <span class="badge object-type-map"></span> Dimensions

A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.

## Definition

```python
# A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
Dimensions: typing.TypeAlias = dict[str, typing.Union[str, list[str]]]
```
