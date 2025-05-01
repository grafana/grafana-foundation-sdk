---
title: <span class="badge object-type-enum"></span> ScaleDistribution
---
# <span class="badge object-type-enum"></span> ScaleDistribution

TODO docs

## Definition

```java
package com.grafana.foundation.common.ScaleDistribution;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ScaleDistribution {
    LINEAR("linear"),
    LOG("log"),
    ORDINAL("ordinal"),
    SYMLOG("symlog"),
    _EMPTY("");

    private final String value;

    private ScaleDistribution(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
