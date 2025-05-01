---
title: <span class="badge object-type-enum"></span> OrderByDirection
---
# <span class="badge object-type-enum"></span> OrderByDirection

## Definition

```java
package com.grafana.foundation.bigquery.OrderByDirection;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum OrderByDirection {
    ASC("ASC"),
    DESC("DESC"),
    _EMPTY("");

    private final String value;

    private OrderByDirection(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
