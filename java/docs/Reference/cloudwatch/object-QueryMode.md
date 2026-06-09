---
title: <span class="badge object-type-enum"></span> QueryMode
---
# <span class="badge object-type-enum"></span> QueryMode

## Definition

```java
package com.grafana.foundation.cloudwatch.QueryMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryMode {
    METRICS("Metrics"),
    LOGS("Logs"),
    ANNOTATIONS("Annotations"),
    _EMPTY("");

    private final String value;

    private QueryMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
