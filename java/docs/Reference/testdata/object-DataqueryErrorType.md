---
title: <span class="badge object-type-enum"></span> DataqueryErrorType
---
# <span class="badge object-type-enum"></span> DataqueryErrorType

## Definition

```java
package com.grafana.foundation.testdata.DataqueryErrorType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DataqueryErrorType {
    SERVER_PANIC("server_panic"),
    FRONTEND_EXCEPTION("frontend_exception"),
    FRONTEND_OBSERVABLE("frontend_observable"),
    _EMPTY("");

    private final String value;

    private DataqueryErrorType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
