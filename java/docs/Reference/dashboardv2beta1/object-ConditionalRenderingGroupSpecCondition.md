---
title: <span class="badge object-type-enum"></span> ConditionalRenderingGroupSpecCondition
---
# <span class="badge object-type-enum"></span> ConditionalRenderingGroupSpecCondition

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpecCondition;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ConditionalRenderingGroupSpecCondition {
    AND("and"),
    OR("or"),
    _EMPTY("");

    private final String value;

    private ConditionalRenderingGroupSpecCondition(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
