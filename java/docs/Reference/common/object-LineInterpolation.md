---
title: <span class="badge object-type-enum"></span> LineInterpolation
---
# <span class="badge object-type-enum"></span> LineInterpolation

TODO docs

## Definition

```java
package com.grafana.foundation.common.LineInterpolation;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LineInterpolation {
    LINEAR("linear"),
    SMOOTH("smooth"),
    STEP_BEFORE("stepBefore"),
    STEP_AFTER("stepAfter"),
    _EMPTY("");

    private final String value;

    private LineInterpolation(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
