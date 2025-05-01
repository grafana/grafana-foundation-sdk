---
title: <span class="badge object-type-enum"></span> PyroscopeQueryType
---
# <span class="badge object-type-enum"></span> PyroscopeQueryType

## Definition

```java
package com.grafana.foundation.grafanapyroscope.PyroscopeQueryType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PyroscopeQueryType {
    METRICS("metrics"),
    PROFILE("profile"),
    BOTH("both"),
    _EMPTY("");

    private final String value;

    private PyroscopeQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
