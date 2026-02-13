// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statetimeline;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = BoolOrUint32Deserializer.class)
@JsonSerialize(using = BoolOrUint32Serializer.class)
public class BoolOrUint32 {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Boolean bool;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Integer uint32;
    protected BoolOrUint32() {}
    public static BoolOrUint32 createBool(Boolean bool) {
        BoolOrUint32 boolOrUint32 = new BoolOrUint32();
        boolOrUint32.bool = bool;
        return boolOrUint32;
    }
    public static BoolOrUint32 createUint32(Integer uint32) {
        BoolOrUint32 boolOrUint32 = new BoolOrUint32();
        boolOrUint32.uint32 = uint32;
        return boolOrUint32;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
