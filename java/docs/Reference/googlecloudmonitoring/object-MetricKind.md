---
title: <span class="badge object-type-enum"></span> MetricKind
---
# <span class="badge object-type-enum"></span> MetricKind

## Definition

```java
package com.grafana.foundation.googlecloudmonitoring.MetricKind;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MetricKind {
    METRIC_KIND_UNSPECIFIED("METRIC_KIND_UNSPECIFIED"),
    GAUGE("GAUGE"),
    DELTA("DELTA"),
    CUMULATIVE("CUMULATIVE"),
    _EMPTY("");

    private final String value;

    private MetricKind(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
