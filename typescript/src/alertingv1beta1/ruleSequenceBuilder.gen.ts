// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';
import * as resource from '../resource';

export class RuleSequenceBuilder implements cog.Builder<alertingv1beta1.RuleSequenceSpec> {
    protected readonly internal: alertingv1beta1.RuleSequenceSpec;

    constructor() {
        this.internal = alertingv1beta1.defaultRuleSequenceSpec();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.RuleSequenceSpec {
        return this.internal;
    }

    trigger(trigger: cog.Builder<alertingv1beta1.IntervalTrigger>): this {
        const triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }

    recordingRules(recordingRules: cog.Builder<alertingv1beta1.RuleRef>[]): this {
        const recordingRulesResources = recordingRules.map(builder1 => builder1.build());
        this.internal.recordingRules = recordingRulesResources;
        return this;
    }

    alertRules(alertRules: cog.Builder<alertingv1beta1.RuleRef>[]): this {
        const alertRulesResources = alertRules.map(builder1 => builder1.build());
        this.internal.alertRules = alertRulesResources;
        return this;
    }
}

/**
 * Creates a resource manifest from a RuleSequence.
 */
export function ruleSequenceManifest(name: string,folder: cog.Builder<alertingv1beta1.RuleSequenceSpec>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("rules.alerting.grafana.app/v1beta1");
	builder.kind("RuleSequence");
	builder.metadata(resource.named(name));
	builder.spec(folder.build());

	return builder;
}

