// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = StringOrBoolDeserializer.class)
@JsonSerialize(using = StringOrBoolSerializer.class)
public class StringOrBool {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Boolean bool;
    protected StringOrBool() {}
    public static StringOrBool createString(String string) {
        StringOrBool stringOrBool = new StringOrBool();
        stringOrBool.string = string;
        return stringOrBool;
    }
    public static StringOrBool createBool(Boolean bool) {
        StringOrBool stringOrBool = new StringOrBool();
        stringOrBool.bool = bool;
        return stringOrBool;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
