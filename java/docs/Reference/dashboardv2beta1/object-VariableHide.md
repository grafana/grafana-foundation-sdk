---
title: <span class="badge object-type-enum"></span> VariableHide
---
# <span class="badge object-type-enum"></span> VariableHide

Determine if the variable shows on dashboard

Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.VariableHide;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Determine if the variable shows on dashboard
// Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableHide {
    DONT_HIDE("dontHide"),
    HIDE_LABEL("hideLabel"),
    HIDE_VARIABLE("hideVariable"),
    IN_CONTROLS_MENU("inControlsMenu"),
    _EMPTY("");

    private final String value;

    private VariableHide(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
