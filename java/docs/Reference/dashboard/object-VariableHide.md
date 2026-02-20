---
title: <span class="badge object-type-enum"></span> VariableHide
---
# <span class="badge object-type-enum"></span> VariableHide

Determine if the variable shows on dashboard

Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing), 3 (show under the controls dropdown menu).

## Definition

```java
package com.grafana.foundation.dashboard.VariableHide;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Determine if the variable shows on dashboard
// Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing), 3 (show under the controls dropdown menu).
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableHide {
    DONT_HIDE(0),
    HIDE_LABEL(1),
    HIDE_VARIABLE(2),
    IN_CONTROLS_MENU(3);

    private final Integer value;

    private VariableHide(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}

```
