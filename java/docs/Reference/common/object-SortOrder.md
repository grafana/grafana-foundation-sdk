---
title: <span class="badge object-type-enum"></span> SortOrder
---
# <span class="badge object-type-enum"></span> SortOrder

TODO docs

## Definition

```java
package com.grafana.foundation.common.SortOrder;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SortOrder {
    ASCENDING("asc"),
    DESCENDING("desc"),
    NONE("none"),
    _EMPTY("");

    private final String value;

    private SortOrder(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
