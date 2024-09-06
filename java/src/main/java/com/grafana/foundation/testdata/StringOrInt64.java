// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = StringOrInt64Deserializer.class)
@JsonSerialize(using = StringOrInt64Serializer.class)
public class StringOrInt64 {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Long int64;
    protected StringOrInt64() {}
    public static StringOrInt64 createString(String string) {
        StringOrInt64 stringOrInt64 = new StringOrInt64();
        stringOrInt64.string = string;
        return stringOrInt64;
    }
    public static StringOrInt64 createInt64(Long int64) {
        StringOrInt64 stringOrInt64 = new StringOrInt64();
        stringOrInt64.int64 = int64;
        return stringOrInt64;
    }
    
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
