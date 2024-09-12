// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

@JsonDeserialize(using = StringOrMapDeserializer.class)
@JsonSerialize(using = StringOrMapSerializer.class)
public class StringOrMap {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Map<String, Object> map;
    protected StringOrMap() {}
    public static StringOrMap createString(String string) {
        StringOrMap stringOrMap = new StringOrMap();
        stringOrMap.string = string;
        return stringOrMap;
    }
    public static StringOrMap createMap(Map<String, Object> map) {
        StringOrMap stringOrMap = new StringOrMap();
        stringOrMap.map = map;
        return stringOrMap;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (string != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(string);
        }
        if (map != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(map);
        }
        
        return null;
    }

}
