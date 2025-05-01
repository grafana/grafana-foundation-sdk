---
title: <span class="badge object-type-enum"></span> SeriesMapping
---
# <span class="badge object-type-enum"></span> SeriesMapping

Auto is "table" in the UI

## Definition

```java
package com.grafana.foundation.xychart.SeriesMapping;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Auto is "table" in the UI
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SeriesMapping {
    AUTO("auto"),
    MANUAL("manual"),
    _EMPTY("");

    private final String value;

    private SeriesMapping(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
