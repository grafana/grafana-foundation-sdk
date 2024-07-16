// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = BuiltinRoleRefOrCustomRoleRefDeserializer.class)
public class BuiltinRoleRefOrCustomRoleRef { 
    @JsonUnwrapped
    public BuiltinRoleRef builtinRoleRef; 
    @JsonUnwrapped
    public CustomRoleRef customRoleRef;
    
    public String toJSON() throws JsonProcessingException {
        if (builtinRoleRef != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(builtinRoleRef);
        }
        if (customRoleRef != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(customRoleRef);
        }
        
        return null;
    }

}
