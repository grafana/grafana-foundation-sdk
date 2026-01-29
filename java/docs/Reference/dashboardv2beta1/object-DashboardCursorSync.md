---
title: <span class="badge object-type-enum"></span> DashboardCursorSync
---
# <span class="badge object-type-enum"></span> DashboardCursorSync

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.DashboardCursorSync;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// "Off" for no shared crosshair or tooltip (default).
// "Crosshair" for shared crosshair.
// "Tooltip" for shared crosshair AND shared tooltip.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DashboardCursorSync {
    CROSSHAIR("Crosshair"),
    TOOLTIP("Tooltip"),
    OFF("Off"),
    _EMPTY("");

    private final String value;

    private DashboardCursorSync(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
