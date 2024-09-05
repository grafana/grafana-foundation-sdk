// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

@JsonDeserialize(using = BuiltinRoleRefOrCustomRoleRefDeserializer.class)
public class BuiltinRoleRefOrCustomRoleRef {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected BuiltinRoleRef builtinRoleRef;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected CustomRoleRef customRoleRef;
    protected BuiltinRoleRefOrCustomRoleRef() {}
    public static BuiltinRoleRefOrCustomRoleRef createBuiltinRoleRef(com.grafana.foundation.cog.Builder<BuiltinRoleRef> builtinRoleRef) {
        BuiltinRoleRefOrCustomRoleRef builtinRoleRefOrCustomRoleRef = new BuiltinRoleRefOrCustomRoleRef();
        builtinRoleRefOrCustomRoleRef.builtinRoleRef = builtinRoleRef.build();
        return builtinRoleRefOrCustomRoleRef;
    }
    public static BuiltinRoleRefOrCustomRoleRef createCustomRoleRef(com.grafana.foundation.cog.Builder<CustomRoleRef> customRoleRef) {
        BuiltinRoleRefOrCustomRoleRef builtinRoleRefOrCustomRoleRef = new BuiltinRoleRefOrCustomRoleRef();
        builtinRoleRefOrCustomRoleRef.customRoleRef = customRoleRef.build();
        return builtinRoleRefOrCustomRoleRef;
    }
    
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
