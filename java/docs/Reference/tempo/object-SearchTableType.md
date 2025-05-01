---
title: <span class="badge object-type-enum"></span> SearchTableType
---
# <span class="badge object-type-enum"></span> SearchTableType

The type of the table that is used to display the search results

## Definition

```java
package com.grafana.foundation.tempo.SearchTableType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// The type of the table that is used to display the search results
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SearchTableType {
    TRACES("traces"),
    SPANS("spans"),
    RAW("raw"),
    _EMPTY("");

    private final String value;

    private SearchTableType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
