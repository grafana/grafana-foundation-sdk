// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class TransformationBuilder implements cog.Builder<dashboardv2.TransformationKind> {
    protected readonly internal: dashboardv2.TransformationKind;

    constructor() {
        this.internal = dashboardv2.defaultTransformationKind();
        this.internal.kind = "Transformation";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.TransformationKind {
        return this.internal;
    }

    // The group is the transformation ID
    group(group: string): this {
        this.internal.group = group;
        return this;
    }

    // Disabled transformations are skipped
    disabled(disabled: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTransformationSpec();
        }
        this.internal.spec.disabled = disabled;
        return this;
    }

    // Optional frame matcher. When missing it will be applied to all results
    filter(filter: dashboardv2.MatcherConfig): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTransformationSpec();
        }
        this.internal.spec.filter = filter;
        return this;
    }

    // Where to pull DataFrames from as input to transformation
    topic(topic: dashboardv2.DataTopic): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTransformationSpec();
        }
        this.internal.spec.topic = topic;
        return this;
    }

    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    options(options: any): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTransformationSpec();
        }
        this.internal.spec.options = options;
        return this;
    }
}

