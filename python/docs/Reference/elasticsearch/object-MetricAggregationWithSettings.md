---
title: <span class="badge object-type-disjunction"></span> MetricAggregationWithSettings
---
# <span class="badge object-type-disjunction"></span> MetricAggregationWithSettings

## Definition

```python
MetricAggregationWithSettings: typing.TypeAlias = typing.Union[elasticsearch.BucketScript, elasticsearch.CumulativeSum, elasticsearch.Derivative, elasticsearch.SerialDiff, elasticsearch.RawData, elasticsearch.RawDocument, elasticsearch.UniqueCount, elasticsearch.Percentiles, elasticsearch.ExtendedStats, elasticsearch.Min, elasticsearch.Max, elasticsearch.Sum, elasticsearch.Average, elasticsearch.MovingAverage, elasticsearch.MovingFunction, elasticsearch.Logs, elasticsearch.Rate, elasticsearch.TopMetrics]
```
