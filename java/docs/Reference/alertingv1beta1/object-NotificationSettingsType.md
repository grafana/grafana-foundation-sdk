---
title: <span class="badge object-type-enum"></span> NotificationSettingsType
---
# <span class="badge object-type-enum"></span> NotificationSettingsType

## Definition

```java
package com.grafana.foundation.alertingv1beta1.NotificationSettingsType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NotificationSettingsType {
    SIMPLIFIED_ROUTING("SimplifiedRouting"),
    NAMED_ROUTING_TREE("NamedRoutingTree"),
    _EMPTY("");

    private final String value;

    private NotificationSettingsType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
