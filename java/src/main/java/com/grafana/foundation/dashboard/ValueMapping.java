// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

// Allow to transform the visual representation of specific data values in a visualization, irrespective of their original units
@JsonDeserialize(using = ValueMappingDeserializer.class)
public class ValueMapping { 
    @JsonUnwrapped
    public ValueMap valueMap; 
    @JsonUnwrapped
    public RangeMap rangeMap; 
    @JsonUnwrapped
    public RegexMap regexMap; 
    @JsonUnwrapped
    public SpecialValueMap specialValueMap;
    
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
