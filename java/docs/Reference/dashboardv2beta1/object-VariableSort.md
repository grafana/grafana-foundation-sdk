---
title: <span class="badge object-type-enum"></span> VariableSort
---
# <span class="badge object-type-enum"></span> VariableSort

Sort variable options

Accepted values are:

`disabled`: No sorting

`alphabeticalAsc`: Alphabetical ASC

`alphabeticalDesc`: Alphabetical DESC

`numericalAsc`: Numerical ASC

`numericalDesc`: Numerical DESC

`alphabeticalCaseInsensitiveAsc`: Alphabetical Case Insensitive ASC

`alphabeticalCaseInsensitiveDesc`: Alphabetical Case Insensitive DESC

`naturalAsc`: Natural ASC

`naturalDesc`: Natural DESC

VariableSort enum with default value

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.VariableSort;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Sort variable options
// Accepted values are:
// `disabled`: No sorting
// `alphabeticalAsc`: Alphabetical ASC
// `alphabeticalDesc`: Alphabetical DESC
// `numericalAsc`: Numerical ASC
// `numericalDesc`: Numerical DESC
// `alphabeticalCaseInsensitiveAsc`: Alphabetical Case Insensitive ASC
// `alphabeticalCaseInsensitiveDesc`: Alphabetical Case Insensitive DESC
// `naturalAsc`: Natural ASC
// `naturalDesc`: Natural DESC
// VariableSort enum with default value
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableSort {
    DISABLED("disabled"),
    ALPHABETICAL_ASC("alphabeticalAsc"),
    ALPHABETICAL_DESC("alphabeticalDesc"),
    NUMERICAL_ASC("numericalAsc"),
    NUMERICAL_DESC("numericalDesc"),
    ALPHABETICAL_CASE_INSENSITIVE_ASC("alphabeticalCaseInsensitiveAsc"),
    ALPHABETICAL_CASE_INSENSITIVE_DESC("alphabeticalCaseInsensitiveDesc"),
    NATURAL_ASC("naturalAsc"),
    NATURAL_DESC("naturalDesc"),
    _EMPTY("");

    private final String value;

    private VariableSort(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
