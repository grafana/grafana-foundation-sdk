// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;

public class AlertRuleSpec {
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("paused")
    public Boolean paused;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("trigger")
    public IntervalTrigger trigger;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public Map<String, String> labels;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("annotations")
    public Map<String, String> annotations;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("for")
    public String forArg;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("keepFiringFor")
    public String keepFiringFor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("missingSeriesEvalsToResolve")
    public Long missingSeriesEvalsToResolve;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("noDataState")
    public NoDataState noDataState;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("execErrState")
    public ExecErrState execErrState;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("notificationSettings")
    public NotificationSettings notificationSettings;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("expressions")
    public Map<String, Expression> expressions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("panelRef")
    public PanelRef panelRef;
    public AlertRuleSpec() {
        this.title = "";
        this.trigger = new com.grafana.foundation.alertingv1beta1.IntervalTrigger();
        this.noDataState = NoDataState.NO_DATA;
        this.execErrState = ExecErrState.ERROR;
        this.expressions = new HashMap<>();
    }
    public AlertRuleSpec(String title,Boolean paused,IntervalTrigger trigger,Map<String, String> labels,Map<String, String> annotations,String forArg,String keepFiringFor,Long missingSeriesEvalsToResolve,NoDataState noDataState,ExecErrState execErrState,NotificationSettings notificationSettings,Map<String, Expression> expressions,PanelRef panelRef) {
        this.title = title;
        this.paused = paused;
        this.trigger = trigger;
        this.labels = labels;
        this.annotations = annotations;
        this.forArg = forArg;
        this.keepFiringFor = keepFiringFor;
        this.missingSeriesEvalsToResolve = missingSeriesEvalsToResolve;
        this.noDataState = noDataState;
        this.execErrState = execErrState;
        this.notificationSettings = notificationSettings;
        this.expressions = expressions;
        this.panelRef = panelRef;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
