---
title: <span class="badge object-type-enum"></span> AnnotationQueryPlacement
---
# <span class="badge object-type-enum"></span> AnnotationQueryPlacement

Annotation Query placement. Defines where the annotation query should be displayed.

- "inControlsMenu" renders the annotation query in the dashboard controls dropdown menu

## Definition

```java
package com.grafana.foundation.dashboard.AnnotationQueryPlacement;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Annotation Query placement. Defines where the annotation query should be displayed.
// - "inControlsMenu" renders the annotation query in the dashboard controls dropdown menu
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AnnotationQueryPlacement {
    IN_CONTROLS_MENU("inControlsMenu"),
    _EMPTY("");

    private final String value;

    private AnnotationQueryPlacement(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
