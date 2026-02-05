// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TransformationBuilder implements cog.Builder<dashboardv2beta1.TransformationKind> {
    protected readonly internal: dashboardv2beta1.TransformationKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultTransformationKind();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TransformationKind {
        return this.internal;
    }

    // The kind of a TransformationKind is the transformation ID
    kind(kind: string): this {
        this.internal.kind = kind;
        return this;
    }

    // Unique identifier of transformer
    id(id: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.id = id;
        return this;
    }

    // Disabled transformations are skipped
    disabled(disabled: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.disabled = disabled;
        return this;
    }

    // Optional frame matcher. When missing it will be applied to all results
    filter(filter: dashboardv2beta1.MatcherConfig): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.filter = filter;
        return this;
    }

    // Where to pull DataFrames from as input to transformation
    topic(topic: dashboardv2beta1.DataTopic): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.topic = topic;
        return this;
    }

    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    options(options: any): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.options = options;
        return this;
    }
}

