---
title: <span class="badge object-type-enum"></span> CloudWatchQueryMode
---
# <span class="badge object-type-enum"></span> CloudWatchQueryMode

## Definition

```java
package com.grafana.foundation.cloudwatch.CloudWatchQueryMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum CloudWatchQueryMode {
    METRICS("Metrics"),
    LOGS("Logs"),
    ANNOTATIONS("Annotations"),
    _EMPTY("");

    private final String value;

    private CloudWatchQueryMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
