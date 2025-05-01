---
title: <span class="badge object-type-enum"></span> MapCenterID
---
# <span class="badge object-type-enum"></span> MapCenterID

## Definition

```java
package com.grafana.foundation.geomap.MapCenterID;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MapCenterID {
    ZERO("zero"),
    COORDS("coords"),
    FIT("fit"),
    _EMPTY("");

    private final String value;

    private MapCenterID(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
