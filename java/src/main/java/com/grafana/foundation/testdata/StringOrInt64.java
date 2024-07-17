// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = StringOrInt64Deserializer.class)
@JsonSerialize(using = StringOrInt64Serializer.class)
public class StringOrInt64 { 
    @JsonUnwrapped
    public String string; 
    @JsonUnwrapped
    public Long int64;
    
    public String toJSON() throws JsonProcessingException {
        if (string != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(string);
        }
        if (int64 != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(int64);
        }
        
        return null;
    }

}
