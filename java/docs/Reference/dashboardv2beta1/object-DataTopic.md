---
title: <span class="badge object-type-enum"></span> DataTopic
---
# <span class="badge object-type-enum"></span> DataTopic

A topic is attached to DataFrame metadata in query results.

This specifies where the data should be used.

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.DataTopic;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// A topic is attached to DataFrame metadata in query results.
// This specifies where the data should be used.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DataTopic {
    SERIES("series"),
    ANNOTATIONS("annotations"),
    ALERT_STATES("alertStates"),
    _EMPTY("");

    private final String value;

    private DataTopic(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
