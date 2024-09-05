// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Allow to transform the visual representation of specific data values in a visualization, irrespective of their original units
@JsonDeserialize(using = ValueMappingDeserializer.class)
public class ValueMapping {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected ValueMap valueMap;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected RangeMap rangeMap;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected RegexMap regexMap;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected SpecialValueMap specialValueMap;
    protected ValueMapping() {}
    public static ValueMapping createValueMap(com.grafana.foundation.cog.Builder<ValueMap> valueMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.valueMap = valueMap.build();
        return valueMapping;
    }
    public static ValueMapping createRangeMap(com.grafana.foundation.cog.Builder<RangeMap> rangeMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.rangeMap = rangeMap.build();
        return valueMapping;
    }
    public static ValueMapping createRegexMap(com.grafana.foundation.cog.Builder<RegexMap> regexMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.regexMap = regexMap.build();
        return valueMapping;
    }
    public static ValueMapping createSpecialValueMap(com.grafana.foundation.cog.Builder<SpecialValueMap> specialValueMap) {
        ValueMapping valueMapping = new ValueMapping();
        valueMapping.specialValueMap = specialValueMap.build();
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
