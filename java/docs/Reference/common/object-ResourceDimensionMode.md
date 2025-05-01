---
title: <span class="badge object-type-enum"></span> ResourceDimensionMode
---
# <span class="badge object-type-enum"></span> ResourceDimensionMode

## Definition

```java
package com.grafana.foundation.common.ResourceDimensionMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ResourceDimensionMode {
    FIXED("fixed"),
    FIELD("field"),
    MAPPING("mapping"),
    _EMPTY("");

    private final String value;

    private ResourceDimensionMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
