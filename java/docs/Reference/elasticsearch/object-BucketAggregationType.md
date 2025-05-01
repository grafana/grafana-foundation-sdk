---
title: <span class="badge object-type-enum"></span> BucketAggregationType
---
# <span class="badge object-type-enum"></span> BucketAggregationType

## Definition

```java
package com.grafana.foundation.elasticsearch.BucketAggregationType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BucketAggregationType {
    TERMS("terms"),
    FILTERS("filters"),
    GEOHASH_GRID("geohash_grid"),
    DATE_HISTOGRAM("date_histogram"),
    HISTOGRAM("histogram"),
    NESTED("nested"),
    _EMPTY("");

    private final String value;

    private BucketAggregationType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
