---
title: <span class="badge object-type-enum"></span> ScaleOrientation
---
# <span class="badge object-type-enum"></span> ScaleOrientation

TODO docs

## Definition

```java
package com.grafana.foundation.common.ScaleOrientation;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ScaleOrientation {
    HORIZONTAL(0),
    VERTICAL(1);

    private final Integer value;

    private ScaleOrientation(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}

```
