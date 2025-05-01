---
title: <span class="badge object-type-enum"></span> LogsSortOrder
---
# <span class="badge object-type-enum"></span> LogsSortOrder

## Definition

```java
package com.grafana.foundation.common.LogsSortOrder;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LogsSortOrder {
    DESCENDING("Descending"),
    ASCENDING("Ascending"),
    _EMPTY("");

    private final String value;

    private LogsSortOrder(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
