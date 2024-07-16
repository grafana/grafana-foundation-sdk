// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Supported value mapping types
// `value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
// `range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
// `regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
// `special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MappingType {
    VALUE_TO_TEXT("value"),
    RANGE_TO_TEXT("range"),
    REGEX_TO_TEXT("regex"),
    SPECIAL_VALUE("special"),
    _EMPTY("");

    private final String value;

    private MappingType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
