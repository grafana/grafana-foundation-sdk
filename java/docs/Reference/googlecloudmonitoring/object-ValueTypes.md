---
title: <span class="badge object-type-enum"></span> ValueTypes
---
# <span class="badge object-type-enum"></span> ValueTypes

## Definition

```java
package com.grafana.foundation.googlecloudmonitoring.ValueTypes;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ValueTypes {
    VALUE_TYPE_UNSPECIFIED("VALUE_TYPE_UNSPECIFIED"),
    BOOL("BOOL"),
    INT64("INT64"),
    DOUBLE("DOUBLE"),
    STRING("STRING"),
    DISTRIBUTION("DISTRIBUTION"),
    MONEY("MONEY"),
    _EMPTY("");

    private final String value;

    private ValueTypes(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
