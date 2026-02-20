---
title: <span class="badge object-type-enum"></span> HttpRequestMethod
---
# <span class="badge object-type-enum"></span> HttpRequestMethod

## Definition

```java
package com.grafana.foundation.dashboard.HttpRequestMethod;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HttpRequestMethod {
    GET("GET"),
    PUT("PUT"),
    POST("POST"),
    DELETE("DELETE"),
    PATCH("PATCH"),
    _EMPTY("");

    private final String value;

    private HttpRequestMethod(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
