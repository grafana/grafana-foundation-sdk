---
title: <span class="badge object-type-enum"></span> StackingMode
---
# <span class="badge object-type-enum"></span> StackingMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.StackingMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum StackingMode {
    NONE("none"),
    NORMAL("normal"),
    PERCENT("percent"),
    _EMPTY("");

    private final String value;

    private StackingMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
