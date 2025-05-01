---
title: <span class="badge object-type-enum"></span> FormatOptions
---
# <span class="badge object-type-enum"></span> FormatOptions

## Definition

```java
package com.grafana.foundation.athena.FormatOptions;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum FormatOptions {
    TIME_SERIES(0),
    TABLE(1),
    LOGS(2);

    private final Integer value;

    private FormatOptions(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}

```
