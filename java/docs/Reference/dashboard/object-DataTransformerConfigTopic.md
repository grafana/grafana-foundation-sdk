---
title: <span class="badge object-type-enum"></span> DataTransformerConfigTopic
---
# <span class="badge object-type-enum"></span> DataTransformerConfigTopic

## Definition

```java
package com.grafana.foundation.dashboard.DataTransformerConfigTopic;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DataTransformerConfigTopic {
    SERIES("series"),
    ANNOTATIONS("annotations"),
    ALERT_STATES("alertStates"),
    _EMPTY("");

    private final String value;

    private DataTransformerConfigTopic(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
