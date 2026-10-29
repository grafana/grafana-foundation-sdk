// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';
import * as resource from '../resource';

export class RecordingRuleBuilder implements cog.Builder<alertingv0alpha1.RecordingRuleSpec> {
    protected readonly internal: alertingv0alpha1.RecordingRuleSpec;

    constructor() {
        this.internal = alertingv0alpha1.defaultRecordingRuleSpec();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.RecordingRuleSpec {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    paused(paused: boolean): this {
        this.internal.paused = paused;
        return this;
    }

    trigger(trigger: cog.Builder<alertingv0alpha1.IntervalTrigger>): this {
        const triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }

    labels(labels: Record<string, alertingv0alpha1.TemplateString>): this {
        this.internal.labels = labels;
        return this;
    }

    metric(metric: alertingv0alpha1.MetricName): this {
        this.internal.metric = metric;
        return this;
    }

    expressions(expressions: alertingv0alpha1.ExpressionMap): this {
        this.internal.expressions = expressions;
        return this;
    }

    targetDatasourceUID(targetDatasourceUID: alertingv0alpha1.DatasourceUID): this {
        this.internal.targetDatasourceUID = targetDatasourceUID;
        return this;
    }
}

/**
 * Creates a resource manifest from a RecordingRule.
 */
export function recordingRuleManifest(name: string,folder: cog.Builder<alertingv0alpha1.RecordingRuleSpec>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("rules.alerting.grafana.app/v0alpha1");
	builder.kind("RecordingRule");
	builder.metadata(resource.named(name));
	builder.spec(folder.build());

	return builder;
}

