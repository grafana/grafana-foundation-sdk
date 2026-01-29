// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class AutoGridLayoutItemBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutItemKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutItemKind();
        this.internal.kind = "AutoGridLayoutItem";
        this.internal.spec.element.kind = "ElementReference";
        this.internal.spec.element.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutItemKind {
        return this.internal;
    }

    element(element: cog.Builder<dashboardv2beta1.ElementReference>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        const elementResource = element.build();
        this.internal.spec.element = elementResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2beta1.AutoGridRepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        if (!this.internal.spec.element) {
            this.internal.spec.element = dashboardv2beta1.defaultElementReference();
        }
        this.internal.spec.element.name = name;
        return this;
    }
}

