---
title: <span class="badge object-type-enum"></span> DataqueryErrorSource
---
# <span class="badge object-type-enum"></span> DataqueryErrorSource

## Definition

```java
package com.grafana.foundation.testdata.DataqueryErrorSource;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DataqueryErrorSource {
    PLUGIN("plugin"),
    DOWNSTREAM("downstream"),
    _EMPTY("");

    private final String value;

    private DataqueryErrorSource(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
