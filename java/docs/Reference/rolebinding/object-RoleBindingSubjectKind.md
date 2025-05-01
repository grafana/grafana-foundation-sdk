---
title: <span class="badge object-type-enum"></span> RoleBindingSubjectKind
---
# <span class="badge object-type-enum"></span> RoleBindingSubjectKind

## Definition

```java
package com.grafana.foundation.rolebinding.RoleBindingSubjectKind;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RoleBindingSubjectKind {
    TEAM("Team"),
    USER("User"),
    _EMPTY("");

    private final String value;

    private RoleBindingSubjectKind(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
