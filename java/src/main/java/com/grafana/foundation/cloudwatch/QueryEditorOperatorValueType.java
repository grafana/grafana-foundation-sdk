// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

@JsonDeserialize(using = QueryEditorOperatorValueTypeDeserializer.class)
@JsonSerialize(using = QueryEditorOperatorValueTypeSerializer.class)
public class QueryEditorOperatorValueType {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Boolean bool;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Long int64;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected List<QueryEditorOperatorType> arrayOfQueryEditorOperatorType;
    protected QueryEditorOperatorValueType() {}
    public static QueryEditorOperatorValueType createString(String string) {
        QueryEditorOperatorValueType queryEditorOperatorValueType = new QueryEditorOperatorValueType();
        queryEditorOperatorValueType.string = string;
        return queryEditorOperatorValueType;
    }
    public static QueryEditorOperatorValueType createBool(Boolean bool) {
        QueryEditorOperatorValueType queryEditorOperatorValueType = new QueryEditorOperatorValueType();
        queryEditorOperatorValueType.bool = bool;
        return queryEditorOperatorValueType;
    }
    public static QueryEditorOperatorValueType createInt64(Long int64) {
        QueryEditorOperatorValueType queryEditorOperatorValueType = new QueryEditorOperatorValueType();
        queryEditorOperatorValueType.int64 = int64;
        return queryEditorOperatorValueType;
    }
    public static QueryEditorOperatorValueType createArrayOfQueryEditorOperatorType(List<QueryEditorOperatorType> arrayOfQueryEditorOperatorType) {
        QueryEditorOperatorValueType queryEditorOperatorValueType = new QueryEditorOperatorValueType();
        queryEditorOperatorValueType.arrayOfQueryEditorOperatorType = arrayOfQueryEditorOperatorType;
        return queryEditorOperatorValueType;
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
        if (arrayOfQueryEditorOperatorType != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(arrayOfQueryEditorOperatorType);
        }
        
        return null;
    }

}
