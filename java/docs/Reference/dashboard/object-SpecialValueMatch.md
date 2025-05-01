---
title: <span class="badge object-type-enum"></span> SpecialValueMatch
---
# <span class="badge object-type-enum"></span> SpecialValueMatch

Special value types supported by the `SpecialValueMap`

## Definition

```java
package com.grafana.foundation.dashboard.SpecialValueMatch;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Special value types supported by the `SpecialValueMap`
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SpecialValueMatch {
    TRUE("true"),
    FALSE("false"),
    NULL("null"),
    NA_N("nan"),
    NULL_AND_NAN("null+nan"),
    EMPTY("empty"),
    _EMPTY("");

    private final String value;

    private SpecialValueMatch(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
