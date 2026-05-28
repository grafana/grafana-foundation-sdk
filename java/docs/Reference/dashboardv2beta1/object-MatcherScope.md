---
title: <span class="badge object-type-enum"></span> MatcherScope
---
# <span class="badge object-type-enum"></span> MatcherScope

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.MatcherScope;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MatcherScope {
    SERIES("series"),
    NESTED("nested"),
    ANNOTATION("annotation"),
    EXEMPLAR("exemplar"),
    _EMPTY("");

    private final String value;

    private MatcherScope(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
