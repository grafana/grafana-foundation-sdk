---
title: <span class="badge object-type-enum"></span> VizOrientation
---
# <span class="badge object-type-enum"></span> VizOrientation

TODO docs

## Definition

```java
package com.grafana.foundation.common.VizOrientation;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VizOrientation {
    AUTO("auto"),
    VERTICAL("vertical"),
    HORIZONTAL("horizontal"),
    _EMPTY("");

    private final String value;

    private VizOrientation(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
