---
title: <span class="badge object-type-enum"></span> BarAlignment
---
# <span class="badge object-type-enum"></span> BarAlignment

TODO docs

## Definition

```java
package com.grafana.foundation.common.BarAlignment;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BarAlignment {
    BEFORE(-1),
    CENTER(0),
    AFTER(1);

    private final Integer value;

    private BarAlignment(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}

```
