---
title: <span class="badge object-type-enum"></span> GraphGradientMode
---
# <span class="badge object-type-enum"></span> GraphGradientMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.GraphGradientMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum GraphGradientMode {
    NONE("none"),
    OPACITY("opacity"),
    HUE("hue"),
    SCHEME("scheme"),
    _EMPTY("");

    private final String value;

    private GraphGradientMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
