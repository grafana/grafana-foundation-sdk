---
title: <span class="badge object-type-enum"></span> LegendDisplayMode
---
# <span class="badge object-type-enum"></span> LegendDisplayMode

TODO docs

Note: "hidden" needs to remain as an option for plugins compatibility

## Definition

```java
package com.grafana.foundation.common.LegendDisplayMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
// Note: "hidden" needs to remain as an option for plugins compatibility
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LegendDisplayMode {
    LIST("list"),
    TABLE("table"),
    HIDDEN("hidden"),
    _EMPTY("");

    private final String value;

    private LegendDisplayMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
