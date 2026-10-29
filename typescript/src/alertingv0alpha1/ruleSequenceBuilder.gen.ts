// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';
import * as resource from '../resource';

export class RuleSequenceBuilder implements cog.Builder<alertingv0alpha1.RuleSequenceSpec> {
    protected readonly internal: alertingv0alpha1.RuleSequenceSpec;

    constructor() {
        this.internal = alertingv0alpha1.defaultRuleSequenceSpec();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.RuleSequenceSpec {
        return this.internal;
    }

    trigger(trigger: cog.Builder<alertingv0alpha1.IntervalTrigger>): this {
        const triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }

    recordingRules(recordingRules: cog.Builder<alertingv0alpha1.RuleRef>[]): this {
        const recordingRulesResources = recordingRules.map(builder1 => builder1.build());
        this.internal.recordingRules = recordingRulesResources;
        return this;
    }

    alertingRules(alertingRules: cog.Builder<alertingv0alpha1.RuleRef>[]): this {
        const alertingRulesResources = alertingRules.map(builder1 => builder1.build());
        this.internal.alertingRules = alertingRulesResources;
        return this;
    }
}

/**
 * Creates a resource manifest from a RuleSequence.
 */
export function ruleSequenceManifest(name: string,folder: cog.Builder<alertingv0alpha1.RuleSequenceSpec>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("rules.alerting.grafana.app/v0alpha1");
	builder.kind("RuleSequence");
	builder.metadata(resource.named(name));
	builder.spec(folder.build());

	return builder;
}

