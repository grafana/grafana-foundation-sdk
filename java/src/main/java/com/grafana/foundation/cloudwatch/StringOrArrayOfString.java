// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.util.List;

@JsonDeserialize(using = StringOrArrayOfStringDeserializer.class)
@JsonSerialize(using = StringOrArrayOfStringSerializer.class)
public class StringOrArrayOfString {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected String string;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected List<String> arrayOfString;
    protected StringOrArrayOfString() {}
    public static StringOrArrayOfString createString(String string) {
        StringOrArrayOfString stringOrArrayOfString = new StringOrArrayOfString();
        stringOrArrayOfString.string = string;
        return stringOrArrayOfString;
    }
    public static StringOrArrayOfString createArrayOfString(List<String> arrayOfString) {
        StringOrArrayOfString stringOrArrayOfString = new StringOrArrayOfString();
        stringOrArrayOfString.arrayOfString = arrayOfString;
        return stringOrArrayOfString;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (string != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(string);
        }
        if (arrayOfString != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(arrayOfString);
        }
        
        return null;
    }

}
