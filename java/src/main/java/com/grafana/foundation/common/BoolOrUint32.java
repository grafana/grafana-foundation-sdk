// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = BoolOrUint32Deserializer.class)
@JsonSerialize(using = BoolOrUint32Serializer.class)
public class BoolOrUint32 { 
    @JsonUnwrapped
    public Boolean bool; 
    @JsonUnwrapped
    public Integer uint32;
    
    public String toJSON() throws JsonProcessingException {
        if (bool != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(bool);
        }
        if (uint32 != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(uint32);
        }
        
        return null;
    }

}
