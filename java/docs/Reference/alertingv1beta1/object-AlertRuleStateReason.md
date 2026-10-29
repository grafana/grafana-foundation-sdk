---
title: <span class="badge object-type-enum"></span> AlertRuleStateReason
---
# <span class="badge object-type-enum"></span> AlertRuleStateReason

## Definition

```java
package com.grafana.foundation.alertingv1beta1.AlertRuleStateReason;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AlertRuleStateReason {
    EVALUATED("Evaluated"),
    KEEP_LAST("KeepLast"),
    _EMPTY("");

    private final String value;

    private AlertRuleStateReason(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
