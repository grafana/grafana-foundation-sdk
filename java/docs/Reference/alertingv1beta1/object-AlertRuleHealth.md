---
title: <span class="badge object-type-enum"></span> AlertRuleHealth
---
# <span class="badge object-type-enum"></span> AlertRuleHealth

## Definition

```java
package com.grafana.foundation.alertingv1beta1.AlertRuleHealth;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AlertRuleHealth {
    UNKNOWN("Unknown"),
    OK("OK"),
    PAUSED("Paused"),
    ERROR("Error"),
    NO_DATA("NoData"),
    _EMPTY("");

    private final String value;

    private AlertRuleHealth(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
