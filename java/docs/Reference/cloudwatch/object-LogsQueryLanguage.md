---
title: <span class="badge object-type-enum"></span> LogsQueryLanguage
---
# <span class="badge object-type-enum"></span> LogsQueryLanguage

## Definition

```java
package com.grafana.foundation.cloudwatch.LogsQueryLanguage;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LogsQueryLanguage {
    CWLI("CWLI"),
    SQL("SQL"),
    PPL("PPL"),
    _EMPTY("");

    private final String value;

    private LogsQueryLanguage(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
