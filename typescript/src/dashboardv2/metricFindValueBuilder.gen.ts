// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Define the MetricFindValue type
export class MetricFindValueBuilder implements cog.Builder<dashboardv2.MetricFindValue> {
    protected readonly internal: dashboardv2.MetricFindValue;

    constructor() {
        this.internal = dashboardv2.defaultMetricFindValue();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.MetricFindValue {
        return this.internal;
    }

    text(text: string): this {
        this.internal.text = text;
        return this;
    }

    value(value: string | number): this {
        this.internal.value = value;
        return this;
    }

    group(group: string): this {
        this.internal.group = group;
        return this;
    }

    expandable(expandable: boolean): this {
        this.internal.expandable = expandable;
        return this;
    }
}

