---
title: <span class="badge object-type-enum"></span> SupportingQueryType
---
# <span class="badge object-type-enum"></span> SupportingQueryType

## Definition

```java
package com.grafana.foundation.loki.SupportingQueryType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SupportingQueryType {
    LOGS_VOLUME("logsVolume"),
    LOGS_SAMPLE("logsSample"),
    DATA_SAMPLE("dataSample"),
    INFINITE_SCROLL("infiniteScroll"),
    _EMPTY("");

    private final String value;

    private SupportingQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
