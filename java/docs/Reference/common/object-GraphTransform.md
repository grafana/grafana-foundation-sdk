---
title: <span class="badge object-type-enum"></span> GraphTransform
---
# <span class="badge object-type-enum"></span> GraphTransform

TODO docs

## Definition

```java
package com.grafana.foundation.common.GraphTransform;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum GraphTransform {
    CONSTANT("constant"),
    NEGATIVE_Y("negative-Y"),
    _EMPTY("");

    private final String value;

    private GraphTransform(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
