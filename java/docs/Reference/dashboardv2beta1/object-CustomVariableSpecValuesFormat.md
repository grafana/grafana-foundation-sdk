---
title: <span class="badge object-type-enum"></span> CustomVariableSpecValuesFormat
---
# <span class="badge object-type-enum"></span> CustomVariableSpecValuesFormat

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.CustomVariableSpecValuesFormat;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum CustomVariableSpecValuesFormat {
    CSV("csv"),
    JSON("json"),
    _EMPTY("");

    private final String value;

    private CustomVariableSpecValuesFormat(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
