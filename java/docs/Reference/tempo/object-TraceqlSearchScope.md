---
title: <span class="badge object-type-enum"></span> TraceqlSearchScope
---
# <span class="badge object-type-enum"></span> TraceqlSearchScope

static fields are pre-set in the UI, dynamic fields are added by the user

## Definition

```java
package com.grafana.foundation.tempo.TraceqlSearchScope;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// static fields are pre-set in the UI, dynamic fields are added by the user
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TraceqlSearchScope {
    INTRINSIC("intrinsic"),
    UNSCOPED("unscoped"),
    RESOURCE("resource"),
    SPAN("span"),
    _EMPTY("");

    private final String value;

    private TraceqlSearchScope(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
