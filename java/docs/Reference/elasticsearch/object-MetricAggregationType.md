---
title: <span class="badge object-type-enum"></span> MetricAggregationType
---
# <span class="badge object-type-enum"></span> MetricAggregationType

## Definition

```java
package com.grafana.foundation.elasticsearch.MetricAggregationType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MetricAggregationType {
    COUNT("count"),
    AVG("avg"),
    SUM("sum"),
    MIN("min"),
    MAX("max"),
    EXTENDED_STATS("extended_stats"),
    PERCENTILES("percentiles"),
    CARDINALITY("cardinality"),
    RAW_DOCUMENT("raw_document"),
    RAW_DATA("raw_data"),
    LOGS("logs"),
    RATE("rate"),
    TOP_METRICS("top_metrics"),
    MOVING_AVG("moving_avg"),
    MOVING_FN("moving_fn"),
    DERIVATIVE("derivative"),
    SERIAL_DIFF("serial_diff"),
    CUMULATIVE_SUM("cumulative_sum"),
    BUCKET_SCRIPT("bucket_script"),
    _EMPTY("");

    private final String value;

    private MetricAggregationType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
