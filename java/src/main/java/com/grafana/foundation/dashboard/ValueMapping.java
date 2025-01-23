// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

// Allow to transform the visual representation of specific data values in a visualization, irrespective of their original units
@JsonDeserialize(using = ValueMappingDeserializer.class)
public class ValueMapping {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ValueMap valueMap;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected RangeMap rangeMap;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected RegexMap regexMap;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected SpecialValueMap specialValueMap;
    protected ValueMapping() {}
    public static ValueMapping createValueMap(ValueMap valueMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.valueMap = valueMap;
        return valueMapping;
    }
    public static ValueMapping createRangeMap(RangeMap rangeMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.rangeMap = rangeMap;
        return valueMapping;
    }
    public static ValueMapping createRegexMap(RegexMap regexMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.regexMap = regexMap;
        return valueMapping;
    }
    public static ValueMapping createSpecialValueMap(SpecialValueMap specialValueMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.specialValueMap = specialValueMap;
        return valueMapping;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (valueMap != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(valueMap);
        }
        if (rangeMap != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(rangeMap);
        }
        if (regexMap != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(regexMap);
        }
        if (specialValueMap != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(specialValueMap);
        }
        
        return null;
    }

}
