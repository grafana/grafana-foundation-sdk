---
title: <span class="badge object-type-enum"></span> TableCellHeight
---
# <span class="badge object-type-enum"></span> TableCellHeight

Height of a table cell

## Definition

```java
package com.grafana.foundation.common.TableCellHeight;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Height of a table cell
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TableCellHeight {
    SM("sm"),
    MD("md"),
    LG("lg"),
    _EMPTY("");

    private final String value;

    private TableCellHeight(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
