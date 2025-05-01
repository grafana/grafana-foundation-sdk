---
title: <span class="badge object-type-enum"></span> MatchType
---
# <span class="badge object-type-enum"></span> MatchType

## Definition

```java
package com.grafana.foundation.alerting.MatchType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MatchType {
    EQUAL("="),
    NOT_EQUAL("!="),
    EQUAL_REGEX("=~"),
    NOT_EQUAL_REGEX("!~"),
    _EMPTY("");

    private final String value;

    private MatchType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
