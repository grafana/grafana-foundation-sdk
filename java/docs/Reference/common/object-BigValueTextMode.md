---
title: <span class="badge object-type-enum"></span> BigValueTextMode
---
# <span class="badge object-type-enum"></span> BigValueTextMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.BigValueTextMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BigValueTextMode {
    AUTO("auto"),
    VALUE("value"),
    VALUE_AND_NAME("value_and_name"),
    NAME("name"),
    NONE("none"),
    _EMPTY("");

    private final String value;

    private BigValueTextMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
