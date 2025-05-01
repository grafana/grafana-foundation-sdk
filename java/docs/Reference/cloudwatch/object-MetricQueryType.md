---
title: <span class="badge object-type-enum"></span> MetricQueryType
---
# <span class="badge object-type-enum"></span> MetricQueryType

## Definition

```java
package com.grafana.foundation.cloudwatch.MetricQueryType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MetricQueryType {
    SEARCH(0),
    QUERY(1);

    private final Integer value;

    private MetricQueryType(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}

```
