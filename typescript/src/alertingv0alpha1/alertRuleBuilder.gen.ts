// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';
import * as resource from '../resource';

export class AlertRuleBuilder implements cog.Builder<alertingv0alpha1.AlertRuleSpec> {
    protected readonly internal: alertingv0alpha1.AlertRuleSpec;

    constructor() {
        this.internal = alertingv0alpha1.defaultAlertRuleSpec();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.AlertRuleSpec {
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

    annotations(annotations: Record<string, alertingv0alpha1.TemplateString>): this {
        this.internal.annotations = annotations;
        return this;
    }

    forVal(forVal: string): this {
        this.internal.for = forVal;
        return this;
    }

    keepFiringFor(keepFiringFor: string): this {
        this.internal.keepFiringFor = keepFiringFor;
        return this;
    }

    missingSeriesEvalsToResolve(missingSeriesEvalsToResolve: number): this {
        if (!(missingSeriesEvalsToResolve >= 0)) {
            throw new Error("missingSeriesEvalsToResolve must be >= 0");
        }
        this.internal.missingSeriesEvalsToResolve = missingSeriesEvalsToResolve;
        return this;
    }

    noDataState(noDataState: alertingv0alpha1.NoDataState): this {
        this.internal.noDataState = noDataState;
        return this;
    }

    execErrState(execErrState: alertingv0alpha1.ExecErrState): this {
        this.internal.execErrState = execErrState;
        return this;
    }

    notificationSettings(notificationSettings: cog.Builder<alertingv0alpha1.NotificationSettings>): this {
        const notificationSettingsResource = notificationSettings.build();
        this.internal.notificationSettings = notificationSettingsResource;
        return this;
    }

    expressions(expressions: alertingv0alpha1.ExpressionMap): this {
        this.internal.expressions = expressions;
        return this;
    }

    panelRef(panelRef: cog.Builder<alertingv0alpha1.PanelRef>): this {
        const panelRefResource = panelRef.build();
        this.internal.panelRef = panelRefResource;
        return this;
    }
}

/**
 * Creates a resource manifest from an AlertRule.
 */
export function alertRuleManifest(name: string,folder: cog.Builder<alertingv0alpha1.AlertRuleSpec>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("rules.alerting.grafana.app/v0alpha1");
	builder.kind("AlertRule");
	builder.metadata(resource.named(name));
	builder.spec(folder.build());

	return builder;
}

