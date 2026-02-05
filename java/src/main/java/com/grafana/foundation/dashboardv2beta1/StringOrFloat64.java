// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = StringOrFloat64Deserializer.class)
@JsonSerialize(using = StringOrFloat64Serializer.class)
public class StringOrFloat64 {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Double float64;
    protected StringOrFloat64() {}
    public static StringOrFloat64 createString(String string) {
        StringOrFloat64 stringOrFloat64 = new StringOrFloat64();
        stringOrFloat64.string = string;
        return stringOrFloat64;
    }
    public static StringOrFloat64 createFloat64(Double float64) {
        StringOrFloat64 stringOrFloat64 = new StringOrFloat64();
        stringOrFloat64.float64 = float64;
        return stringOrFloat64;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
