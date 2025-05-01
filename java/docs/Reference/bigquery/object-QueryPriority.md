---
title: <span class="badge object-type-enum"></span> QueryPriority
---
# <span class="badge object-type-enum"></span> QueryPriority

## Definition

```java
package com.grafana.foundation.bigquery.QueryPriority;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryPriority {
    INTERACTIVE("INTERACTIVE"),
    BATCH("BATCH"),
    _EMPTY("");

    private final String value;

    private QueryPriority(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
