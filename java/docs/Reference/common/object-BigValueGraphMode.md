---
title: <span class="badge object-type-enum"></span> BigValueGraphMode
---
# <span class="badge object-type-enum"></span> BigValueGraphMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.BigValueGraphMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BigValueGraphMode {
    NONE("none"),
    LINE("line"),
    AREA("area"),
    _EMPTY("");

    private final String value;

    private BigValueGraphMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
