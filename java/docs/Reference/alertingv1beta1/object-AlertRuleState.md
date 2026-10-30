---
title: <span class="badge object-type-enum"></span> AlertRuleState
---
# <span class="badge object-type-enum"></span> AlertRuleState

## Definition

```java
package com.grafana.foundation.alertingv1beta1.AlertRuleState;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AlertRuleState {
    INACTIVE("Inactive"),
    HEALTHY("Healthy"),
    FIRING("Firing"),
    PENDING("Pending"),
    RECOVERING("Recovering"),
    _EMPTY("");

    private final String value;

    private AlertRuleState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
