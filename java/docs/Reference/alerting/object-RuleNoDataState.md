---
title: <span class="badge object-type-enum"></span> RuleNoDataState
---
# <span class="badge object-type-enum"></span> RuleNoDataState

## Definition

```java
package com.grafana.foundation.alerting.RuleNoDataState;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RuleNoDataState {
    ALERTING("Alerting"),
    NO_DATA("NoData"),
    OK("OK"),
    _EMPTY("");

    private final String value;

    private RuleNoDataState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
