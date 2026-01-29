// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Define the MetricFindValue type
export class MetricFindValueBuilder implements cog.Builder<dashboardv2beta1.MetricFindValue> {
    protected readonly internal: dashboardv2beta1.MetricFindValue;

    constructor() {
        this.internal = dashboardv2beta1.defaultMetricFindValue();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.MetricFindValue {
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

