---
title: <span class="badge object-type-enum"></span> TermsOrder
---
# <span class="badge object-type-enum"></span> TermsOrder

## Definition

```java
package com.grafana.foundation.elasticsearch.TermsOrder;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TermsOrder {
    DESC("desc"),
    ASC("asc"),
    _EMPTY("");

    private final String value;

    private TermsOrder(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
