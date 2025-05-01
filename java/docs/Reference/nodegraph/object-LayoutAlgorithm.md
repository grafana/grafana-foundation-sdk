---
title: <span class="badge object-type-enum"></span> LayoutAlgorithm
---
# <span class="badge object-type-enum"></span> LayoutAlgorithm

## Definition

```java
package com.grafana.foundation.nodegraph.LayoutAlgorithm;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LayoutAlgorithm {
    LAYERED("layered"),
    FORCE("force"),
    GRID("grid"),
    _EMPTY("");

    private final String value;

    private LayoutAlgorithm(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
