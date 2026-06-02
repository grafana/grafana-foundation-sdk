---
title: <span class="badge object-type-enum"></span> AlignmentTypes
---
# <span class="badge object-type-enum"></span> AlignmentTypes

## Definition

```java
package com.grafana.foundation.googlecloudmonitoring.AlignmentTypes;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AlignmentTypes {
    ALIGNDELTA("ALIGN_DELTA"),
    ALIGNRATE("ALIGN_RATE"),
    ALIGNINTERPOLATE("ALIGN_INTERPOLATE"),
    ALIGNNEXTOLDER("ALIGN_NEXT_OLDER"),
    ALIGNMIN("ALIGN_MIN"),
    ALIGNMAX("ALIGN_MAX"),
    ALIGNMEAN("ALIGN_MEAN"),
    ALIGNCOUNT("ALIGN_COUNT"),
    ALIGNSUM("ALIGN_SUM"),
    ALIGNSTDDEV("ALIGN_STDDEV"),
    ALIGNCOUNTTRUE("ALIGN_COUNT_TRUE"),
    ALIGNCOUNTFALSE("ALIGN_COUNT_FALSE"),
    ALIGNFRACTIONTRUE("ALIGN_FRACTION_TRUE"),
    ALIGNPERCENTILE99("ALIGN_PERCENTILE_99"),
    ALIGNPERCENTILE95("ALIGN_PERCENTILE_95"),
    ALIGNPERCENTILE50("ALIGN_PERCENTILE_50"),
    ALIGNPERCENTILE05("ALIGN_PERCENTILE_05"),
    ALIGNPERCENTCHANGE("ALIGN_PERCENT_CHANGE"),
    ALIGNNONE("ALIGN_NONE"),
    _EMPTY("");

    private final String value;

    private AlignmentTypes(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
