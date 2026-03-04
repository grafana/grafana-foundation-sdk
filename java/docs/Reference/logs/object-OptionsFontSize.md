---
title: <span class="badge object-type-enum"></span> OptionsFontSize
---
# <span class="badge object-type-enum"></span> OptionsFontSize

## Definition

```java
package com.grafana.foundation.logs.OptionsFontSize;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum OptionsFontSize {
    DEFAULT("default"),
    SMALL("small"),
    _EMPTY("");

    private final String value;

    private OptionsFontSize(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
