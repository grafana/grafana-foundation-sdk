// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.grafana.foundation.resource.ManifestBuilder;
import com.grafana.foundation.resource.Resource;
public class Alertingv1beta1 {    
    /**
     * Creates a resource manifest from an AlertRule.
     */
    public static ManifestBuilder alertRuleManifest(String name,com.grafana.foundation.cog.Builder<AlertRuleSpec> folder) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("rules.alerting.grafana.app/v1beta1");
        builder.kind("AlertRule");
        builder.metadata(Resource.named(name));
        builder.spec(folder.build());
        return builder;
    }
    
    /**
     * Creates a resource manifest from a RecordingRule.
     */
    public static ManifestBuilder recordingRuleManifest(String name,com.grafana.foundation.cog.Builder<RecordingRuleSpec> folder) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("rules.alerting.grafana.app/v1beta1");
        builder.kind("RecordingRule");
        builder.metadata(Resource.named(name));
        builder.spec(folder.build());
        return builder;
    }
    
    /**
     * Creates a resource manifest from a RuleSequence.
     */
    public static ManifestBuilder ruleSequenceManifest(String name,com.grafana.foundation.cog.Builder<RuleSequenceSpec> folder) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("rules.alerting.grafana.app/v1beta1");
        builder.kind("RuleSequence");
        builder.metadata(Resource.named(name));
        builder.spec(folder.build());
        return builder;
    }
}
