---
title: <span class="badge object-type-enum"></span> VariableRefresh
---
# <span class="badge object-type-enum"></span> VariableRefresh

Options to config when to refresh a variable

`0`: Never refresh the variable

`1`: Queries the data source every time the dashboard loads.

`2`: Queries the data source when the dashboard time range changes.

## Definition

```java
package com.grafana.foundation.dashboard.VariableRefresh;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Options to config when to refresh a variable
// `0`: Never refresh the variable
// `1`: Queries the data source every time the dashboard loads.
// `2`: Queries the data source when the dashboard time range changes.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableRefresh {
    NEVER(0),
    ON_DASHBOARD_LOAD(1),
    ON_TIME_RANGE_CHANGED(2);

    private final Integer value;

    private VariableRefresh(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}

```
