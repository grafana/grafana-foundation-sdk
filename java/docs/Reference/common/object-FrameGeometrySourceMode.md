---
title: <span class="badge object-type-enum"></span> FrameGeometrySourceMode
---
# <span class="badge object-type-enum"></span> FrameGeometrySourceMode

## Definition

```java
package com.grafana.foundation.common.FrameGeometrySourceMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum FrameGeometrySourceMode {
    AUTO("auto"),
    GEOHASH("geohash"),
    COORDS("coords"),
    LOOKUP("lookup"),
    _EMPTY("");

    private final String value;

    private FrameGeometrySourceMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
