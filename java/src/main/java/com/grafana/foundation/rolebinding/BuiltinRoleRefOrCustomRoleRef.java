// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = BuiltinRoleRefOrCustomRoleRefDeserializer.class)
public class BuiltinRoleRefOrCustomRoleRef {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected BuiltinRoleRef builtinRoleRef;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CustomRoleRef customRoleRef;
    protected BuiltinRoleRefOrCustomRoleRef() {}
    public static BuiltinRoleRefOrCustomRoleRef createBuiltinRoleRef(BuiltinRoleRef builtinRoleRef) {
        BuiltinRoleRefOrCustomRoleRef builtinRoleRefOrCustomRoleRef = new BuiltinRoleRefOrCustomRoleRef();
        builtinRoleRefOrCustomRoleRef.builtinRoleRef = builtinRoleRef;
        return builtinRoleRefOrCustomRoleRef;
    }
    public static BuiltinRoleRefOrCustomRoleRef createCustomRoleRef(CustomRoleRef customRoleRef) {
        BuiltinRoleRefOrCustomRoleRef builtinRoleRefOrCustomRoleRef = new BuiltinRoleRefOrCustomRoleRef();
        builtinRoleRefOrCustomRoleRef.customRoleRef = customRoleRef;
        return builtinRoleRefOrCustomRoleRef;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
