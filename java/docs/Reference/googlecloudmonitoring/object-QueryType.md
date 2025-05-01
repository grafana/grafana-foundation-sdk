---
title: <span class="badge object-type-enum"></span> QueryType
---
# <span class="badge object-type-enum"></span> QueryType

Defines the supported queryTypes.

## Definition

```java
package com.grafana.foundation.googlecloudmonitoring.QueryType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Defines the supported queryTypes.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryType {
    TIME_SERIES_LIST("timeSeriesList"),
    TIME_SERIES_QUERY("timeSeriesQuery"),
    SLO("slo"),
    ANNOTATION("annotation"),
    PROMQL("promQL"),
    _EMPTY("");

    private final String value;

    private QueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
