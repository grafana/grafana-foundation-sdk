---
title: <span class="badge object-type-enum"></span> PreprocessorType
---
# <span class="badge object-type-enum"></span> PreprocessorType

Types of pre-processor available. Defined by the metric.

## Definition

```java
package com.grafana.foundation.googlecloudmonitoring.PreprocessorType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Types of pre-processor available. Defined by the metric.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PreprocessorType {
    NONE("none"),
    RATE("rate"),
    DELTA("delta"),
    _EMPTY("");

    private final String value;

    private PreprocessorType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
