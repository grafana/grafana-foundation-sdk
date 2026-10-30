---
title: <span class="badge object-type-enum"></span> NoDataState
---
# <span class="badge object-type-enum"></span> NoDataState

## Definition

```java
package com.grafana.foundation.alertingv1beta1.NoDataState;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NoDataState {
    NO_DATA("NoData"),
    OK("Ok"),
    ALERTING("Alerting"),
    KEEP_LAST("KeepLast"),
    _EMPTY("");

    private final String value;

    private NoDataState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
