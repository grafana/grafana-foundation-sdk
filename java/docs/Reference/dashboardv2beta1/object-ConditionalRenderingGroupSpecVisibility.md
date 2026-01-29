---
title: <span class="badge object-type-enum"></span> ConditionalRenderingGroupSpecVisibility
---
# <span class="badge object-type-enum"></span> ConditionalRenderingGroupSpecVisibility

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpecVisibility;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ConditionalRenderingGroupSpecVisibility {
    SHOW("show"),
    HIDE("hide"),
    _EMPTY("");

    private final String value;

    private ConditionalRenderingGroupSpecVisibility(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
