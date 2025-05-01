---
title: <span class="badge object-type-enum"></span> MetricsQueryType
---
# <span class="badge object-type-enum"></span> MetricsQueryType

## Definition

```java
package com.grafana.foundation.tempo.MetricsQueryType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MetricsQueryType {
    RANGE("range"),
    INSTANT("instant"),
    _EMPTY("");

    private final String value;

    private MetricsQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
