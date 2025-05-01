---
title: <span class="badge object-type-enum"></span> RoleRefKind
---
# <span class="badge object-type-enum"></span> RoleRefKind

## Definition

```java
package com.grafana.foundation.accesspolicy.RoleRefKind;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RoleRefKind {
    ROLE("Role"),
    BUILTIN_ROLE("BuiltinRole"),
    TEAM("Team"),
    USER("User"),
    _EMPTY("");

    private final String value;

    private RoleRefKind(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
