---
title: <span class="badge object-type-enum"></span> NotebookLayoutItemSpecSource
---
# <span class="badge object-type-enum"></span> NotebookLayoutItemSpecSource

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.NotebookLayoutItemSpecSource;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NotebookLayoutItemSpecSource {
    ASSISTANT("assistant"),
    USER("user"),
    _EMPTY("");

    private final String value;

    private NotebookLayoutItemSpecSource(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
