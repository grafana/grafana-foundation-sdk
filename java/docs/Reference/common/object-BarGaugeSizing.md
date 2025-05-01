---
title: <span class="badge object-type-enum"></span> BarGaugeSizing
---
# <span class="badge object-type-enum"></span> BarGaugeSizing

Allows for the bar gauge size to be set explicitly

## Definition

```java
package com.grafana.foundation.common.BarGaugeSizing;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Allows for the bar gauge size to be set explicitly
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BarGaugeSizing {
    AUTO("auto"),
    MANUAL("manual"),
    _EMPTY("");

    private final String value;

    private BarGaugeSizing(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
