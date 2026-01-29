// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = VariableKindDeserializer.class)
public class VariableKind {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryVariableKind queryVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TextVariableKind textVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ConstantVariableKind constantVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected DatasourceVariableKind datasourceVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected IntervalVariableKind intervalVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CustomVariableKind customVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected GroupByVariableKind groupByVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected AdhocVariableKind adhocVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected SwitchVariableKind switchVariableKind;
    protected VariableKind() {}
    public static VariableKind createQueryVariableKind(QueryVariableKind queryVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.queryVariableKind = queryVariableKind;
        return variableKind;
    }
    public static VariableKind createTextVariableKind(TextVariableKind textVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.textVariableKind = textVariableKind;
        return variableKind;
    }
    public static VariableKind createConstantVariableKind(ConstantVariableKind constantVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.constantVariableKind = constantVariableKind;
        return variableKind;
    }
    public static VariableKind createDatasourceVariableKind(DatasourceVariableKind datasourceVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.datasourceVariableKind = datasourceVariableKind;
        return variableKind;
    }
    public static VariableKind createIntervalVariableKind(IntervalVariableKind intervalVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.intervalVariableKind = intervalVariableKind;
        return variableKind;
    }
    public static VariableKind createCustomVariableKind(CustomVariableKind customVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.customVariableKind = customVariableKind;
        return variableKind;
    }
    public static VariableKind createGroupByVariableKind(GroupByVariableKind groupByVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.groupByVariableKind = groupByVariableKind;
        return variableKind;
    }
    public static VariableKind createAdhocVariableKind(AdhocVariableKind adhocVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.adhocVariableKind = adhocVariableKind;
        return variableKind;
    }
    public static VariableKind createSwitchVariableKind(SwitchVariableKind switchVariableKind) {
        VariableKind variableKind = new VariableKind();
        variableKind.switchVariableKind = switchVariableKind;
        return variableKind;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
