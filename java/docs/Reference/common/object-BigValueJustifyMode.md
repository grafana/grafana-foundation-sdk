---
title: <span class="badge object-type-enum"></span> BigValueJustifyMode
---
# <span class="badge object-type-enum"></span> BigValueJustifyMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.BigValueJustifyMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BigValueJustifyMode {
    AUTO("auto"),
    CENTER("center"),
    _EMPTY("");

    private final String value;

    private BigValueJustifyMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
