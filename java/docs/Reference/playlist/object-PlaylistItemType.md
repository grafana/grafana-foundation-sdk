---
title: <span class="badge object-type-enum"></span> PlaylistItemType
---
# <span class="badge object-type-enum"></span> PlaylistItemType

## Definition

```java
package com.grafana.foundation.playlist.PlaylistItemType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PlaylistItemType {
    DASHBOARD_BY_UID("dashboard_by_uid"),
    DASHBOARD_BY_ID("dashboard_by_id"),
    DASHBOARD_BY_TAG("dashboard_by_tag"),
    _EMPTY("");

    private final String value;

    private PlaylistItemType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
