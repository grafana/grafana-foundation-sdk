---
title: <span class="badge object-type-enum"></span> ConditionalRenderingVariableSpecOperator
---
# <span class="badge object-type-enum"></span> ConditionalRenderingVariableSpecOperator

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.ConditionalRenderingVariableSpecOperator;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ConditionalRenderingVariableSpecOperator {
    EQUALS("equals"),
    NOT_EQUALS("notEquals"),
    MATCHES("matches"),
    NOT_MATCHES("notMatches"),
    _EMPTY("");

    private final String value;

    private ConditionalRenderingVariableSpecOperator(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
