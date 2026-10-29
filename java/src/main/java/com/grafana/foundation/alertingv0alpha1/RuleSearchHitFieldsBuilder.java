// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import java.util.Map;
import java.util.List;

/**
 * #RuleSearchHitFields is the per-kind field payload returned on each hit.
 * It carries the union of alert- and recording-rule search fields; only the
 * fields relevant to a hit's kind are populated. This maps to the kind's
 * declared searchFields.
 */
public class RuleSearchHitFieldsBuilder implements com.grafana.foundation.cog.Builder<RuleSearchHitFields> {
    protected final RuleSearchHitFields internal;
    
    public RuleSearchHitFieldsBuilder() {
        this.internal = new RuleSearchHitFields();
    }
    public RuleSearchHitFieldsBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder folder(String folder) {
        this.internal.folder = folder;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder paused(Boolean paused) {
        this.internal.paused = paused;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder datasourceUIDs(List<String> datasourceUIDs) {
        this.internal.datasourceUIDs = datasourceUIDs;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder annotations(Map<String, String> annotations) {
        this.internal.annotations = annotations;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder forArg(String forArg) {
        this.internal.forArg = forArg;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder keepFiringFor(String keepFiringFor) {
        this.internal.keepFiringFor = keepFiringFor;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder dashboardUID(String dashboardUID) {
        this.internal.dashboardUID = dashboardUID;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder panelID(Long panelID) {
        this.internal.panelID = panelID;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder receiver(String receiver) {
        this.internal.receiver = receiver;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder notificationType(String notificationType) {
        this.internal.notificationType = notificationType;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder routingTree(String routingTree) {
        this.internal.routingTree = routingTree;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder metric(String metric) {
        this.internal.metric = metric;
        return this;
    }
    
    public RuleSearchHitFieldsBuilder targetDatasourceUID(String targetDatasourceUID) {
        this.internal.targetDatasourceUID = targetDatasourceUID;
        return this;
    }
    public RuleSearchHitFields build() {
        return this.internal;
    }
}
