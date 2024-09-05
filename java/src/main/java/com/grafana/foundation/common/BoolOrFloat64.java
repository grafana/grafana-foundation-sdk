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
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Boolean bool;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Double float64;
    protected BoolOrFloat64() {}
    public static BoolOrFloat64 createBool(Boolean bool) {
        BoolOrFloat64 boolOrFloat64 = new BoolOrFloat64();
        boolOrFloat64.bool = bool;
        return boolOrFloat64;
    }
    public static BoolOrFloat64 createFloat64(Double float64) {
        BoolOrFloat64 boolOrFloat64 = new BoolOrFloat64();
        boolOrFloat64.float64 = float64;
        return boolOrFloat64;
    }
    
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
