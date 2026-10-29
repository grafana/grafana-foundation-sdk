// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class ExpressionBuilder implements cog.Builder<alertingv0alpha1.Expression> {
    protected readonly internal: alertingv0alpha1.Expression;

    constructor() {
        this.internal = alertingv0alpha1.defaultExpression();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.Expression {
        return this.internal;
    }

    // The type of query if this is a query expression
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    relativeTimeRange(relativeTimeRange: cog.Builder<alertingv0alpha1.RelativeTimeRange>): this {
        const relativeTimeRangeResource = relativeTimeRange.build();
        this.internal.relativeTimeRange = relativeTimeRangeResource;
        return this;
    }

    // The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
    datasourceUID(datasourceUID: alertingv0alpha1.DatasourceUID): this {
        this.internal.datasourceUID = datasourceUID;
        return this;
    }

    model(model: any): this {
        this.internal.model = model;
        return this;
    }

    // Used to mark the expression to be used as the final source for the rule evaluation
    // Only one expression in a rule can be marked as the source
    // For AlertRules, this is the expression that will be evaluated against the alerting condition
    // For RecordingRules, this is the expression that will be recorded
    source(source: boolean): this {
        this.internal.source = source;
        return this;
    }
}

