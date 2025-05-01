---
title: <span class="badge object-type-enum"></span> NodesQueryType
---
# <span class="badge object-type-enum"></span> NodesQueryType

## Definition

```java
package com.grafana.foundation.testdata.NodesQueryType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NodesQueryType {
    RANDOM("random"),
    RESPONSE("response"),
    RANDOM_EDGES("random edges"),
    _EMPTY("");

    private final String value;

    private NodesQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
