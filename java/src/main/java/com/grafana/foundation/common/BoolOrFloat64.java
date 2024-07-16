// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = BoolOrFloat64Deserializer.class)
@JsonSerialize(using = BoolOrFloat64Serializer.class)
public class BoolOrFloat64 { 
    @JsonUnwrapped
    public Boolean bool; 
    @JsonUnwrapped
    public Double float64;
    
    public String toJSON() throws JsonProcessingException {
        if (bool != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(bool);
        }
        if (float64 != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(float64);
        }
        
        return null;
    }

}
