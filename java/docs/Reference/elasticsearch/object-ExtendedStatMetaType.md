---
title: <span class="badge object-type-enum"></span> ExtendedStatMetaType
---
# <span class="badge object-type-enum"></span> ExtendedStatMetaType

## Definition

```java
package com.grafana.foundation.elasticsearch.ExtendedStatMetaType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ExtendedStatMetaType {
    AVG("avg"),
    MIN("min"),
    MAX("max"),
    SUM("sum"),
    COUNT("count"),
    STD_DEVIATION("std_deviation"),
    STD_DEVIATION_BOUNDS_UPPER("std_deviation_bounds_upper"),
    STD_DEVIATION_BOUNDS_LOWER("std_deviation_bounds_lower"),
    _EMPTY("");

    private final String value;

    private ExtendedStatMetaType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
