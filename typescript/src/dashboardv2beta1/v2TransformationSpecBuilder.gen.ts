// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

/**
 * Dashboard v2 shape: no `id`, it is carried by the parent's `group`.
 */
export class V2TransformationSpecBuilder implements cog.Builder<dashboardv2beta1.V2TransformationSpec> {
    protected readonly internal: dashboardv2beta1.V2TransformationSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultV2TransformationSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.V2TransformationSpec {
        return this.internal;
    }

    // Disabled transformations are skipped
    disabled(disabled: boolean): this {
        this.internal.disabled = disabled;
        return this;
    }

    // Optional frame matcher. When missing it will be applied to all results
    filter(filter: dashboardv2beta1.MatcherConfig): this {
        this.internal.filter = filter;
        return this;
    }

    // Where to pull DataFrames from as input to transformation
    topic(topic: dashboardv2beta1.DataTopic): this {
        this.internal.topic = topic;
        return this;
    }

    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    options(options: any): this {
        this.internal.options = options;
        return this;
    }
}

