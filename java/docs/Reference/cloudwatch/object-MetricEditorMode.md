---
title: <span class="badge object-type-enum"></span> MetricEditorMode
---
# <span class="badge object-type-enum"></span> MetricEditorMode

## Definition

```java
package com.grafana.foundation.cloudwatch.MetricEditorMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MetricEditorMode {
    BUILDER(0),
    CODE(1);

    private final Integer value;

    private MetricEditorMode(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}

```
