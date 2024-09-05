// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = QueryEditorOperatorTypeDeserializer.class)
@JsonSerialize(using = QueryEditorOperatorTypeSerializer.class)
public class QueryEditorOperatorType {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected String string;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Boolean bool;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Long int64;
    protected QueryEditorOperatorType() {}
    public static QueryEditorOperatorType createString(String string) {
        QueryEditorOperatorType queryEditorOperatorType = new QueryEditorOperatorType();
        queryEditorOperatorType.string = string;
        return queryEditorOperatorType;
    }
    public static QueryEditorOperatorType createBool(Boolean bool) {
        QueryEditorOperatorType queryEditorOperatorType = new QueryEditorOperatorType();
        queryEditorOperatorType.bool = bool;
        return queryEditorOperatorType;
    }
    public static QueryEditorOperatorType createInt64(Long int64) {
        QueryEditorOperatorType queryEditorOperatorType = new QueryEditorOperatorType();
        queryEditorOperatorType.int64 = int64;
        return queryEditorOperatorType;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (string != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(string);
        }
        if (bool != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(bool);
        }
        if (int64 != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(int64);
        }
        
        return null;
    }

}
