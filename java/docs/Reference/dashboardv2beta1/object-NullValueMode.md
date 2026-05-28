---
title: <span class="badge object-type-enum"></span> NullValueMode
---
# <span class="badge object-type-enum"></span> NullValueMode

How null values should be handled

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.NullValueMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// How null values should be handled
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NullValueMode {
    NULL("null"),
    CONNECTED("connected"),
    NULL_AS_ZERO("null as zero"),
    _EMPTY("");

    private final String value;

    private NullValueMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
