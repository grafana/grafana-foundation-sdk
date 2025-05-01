---
title: <span class="badge object-type-enum"></span> VisibilityMode
---
# <span class="badge object-type-enum"></span> VisibilityMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.VisibilityMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VisibilityMode {
    AUTO("auto"),
    NEVER("never"),
    ALWAYS("always"),
    _EMPTY("");

    private final String value;

    private VisibilityMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
