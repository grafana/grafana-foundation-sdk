// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

/**
 * #RuleSearchHitFields is the per-kind field payload returned on each hit.
 * It carries the union of alert- and recording-rule search fields; only the
 * fields relevant to a hit's kind are populated. This maps to the kind's
 * declared searchFields.
 */
export class RuleSearchHitFieldsBuilder implements cog.Builder<alertingv0alpha1.RuleSearchHitFields> {
    protected readonly internal: alertingv0alpha1.RuleSearchHitFields;

    constructor() {
        this.internal = alertingv0alpha1.defaultRuleSearchHitFields();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.RuleSearchHitFields {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    folder(folder: string): this {
        this.internal.folder = folder;
        return this;
    }

    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    paused(paused: boolean): this {
        this.internal.paused = paused;
        return this;
    }

    labels(labels: Record<string, string>): this {
        this.internal.labels = labels;
        return this;
    }

    datasourceUIDs(datasourceUIDs: string[]): this {
        this.internal.datasourceUIDs = datasourceUIDs;
        return this;
    }

    // Alert-rule fields.
    annotations(annotations: Record<string, string>): this {
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

    dashboardUID(dashboardUID: string): this {
        this.internal.dashboardUID = dashboardUID;
        return this;
    }

    panelID(panelID: number): this {
        this.internal.panelID = panelID;
        return this;
    }

    receiver(receiver: string): this {
        this.internal.receiver = receiver;
        return this;
    }

    notificationType(notificationType: string): this {
        this.internal.notificationType = notificationType;
        return this;
    }

    routingTree(routingTree: string): this {
        this.internal.routingTree = routingTree;
        return this;
    }

    // Recording-rule fields.
    metric(metric: string): this {
        this.internal.metric = metric;
        return this;
    }

    targetDatasourceUID(targetDatasourceUID: string): this {
        this.internal.targetDatasourceUID = targetDatasourceUID;
        return this;
    }
}

