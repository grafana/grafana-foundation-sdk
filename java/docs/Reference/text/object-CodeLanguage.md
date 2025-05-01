---
title: <span class="badge object-type-enum"></span> CodeLanguage
---
# <span class="badge object-type-enum"></span> CodeLanguage

## Definition

```java
package com.grafana.foundation.text.CodeLanguage;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum CodeLanguage {
    JSON("json"),
    YAML("yaml"),
    XML("xml"),
    TYPESCRIPT("typescript"),
    SQL("sql"),
    GO("go"),
    MARKDOWN("markdown"),
    HTML("html"),
    PLAINTEXT("plaintext"),
    _EMPTY("");

    private final String value;

    private CodeLanguage(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
