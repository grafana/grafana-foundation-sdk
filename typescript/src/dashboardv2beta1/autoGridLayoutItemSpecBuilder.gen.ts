// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class AutoGridLayoutItemSpecBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutItemSpec> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutItemSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutItemSpec {
        return this.internal;
    }

    element(element: cog.Builder<dashboardv2beta1.ElementReference>): this {
        const elementResource = element.build();
        this.internal.element = elementResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2beta1.AutoGridRepeatOptions>): this {
        const repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this {
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }
}

