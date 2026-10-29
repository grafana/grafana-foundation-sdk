---
title: <span class="badge object-type-enum"></span> ExecErrState
---
# <span class="badge object-type-enum"></span> ExecErrState

## Definition

```java
package com.grafana.foundation.alertingv0alpha1.ExecErrState;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ExecErrState {
    ERROR("Error"),
    OK("Ok"),
    ALERTING("Alerting"),
    KEEP_LAST("KeepLast"),
    _EMPTY("");

    private final String value;

    private ExecErrState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
