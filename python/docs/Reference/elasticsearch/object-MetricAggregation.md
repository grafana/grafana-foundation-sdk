---
title: <span class="badge object-type-disjunction"></span> MetricAggregation
---
# <span class="badge object-type-disjunction"></span> MetricAggregation

## Definition

```python
MetricAggregation: typing.TypeAlias = typing.Union[elasticsearch.Count, elasticsearch.MovingAverage, elasticsearch.Derivative, elasticsearch.CumulativeSum, elasticsearch.BucketScript, elasticsearch.SerialDiff, elasticsearch.RawData, elasticsearch.RawDocument, elasticsearch.UniqueCount, elasticsearch.Percentiles, elasticsearch.ExtendedStats, elasticsearch.Min, elasticsearch.Max, elasticsearch.Sum, elasticsearch.Average, elasticsearch.MovingFunction, elasticsearch.Logs, elasticsearch.Rate, elasticsearch.TopMetrics]
```
