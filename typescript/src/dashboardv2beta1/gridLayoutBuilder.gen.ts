// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class GridLayoutBuilder implements cog.Builder<dashboardv2beta1.GridLayoutKind> {
    protected readonly internal: dashboardv2beta1.GridLayoutKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultGridLayoutKind();
        this.internal.kind = "GridLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.GridLayoutKind {
        return this.internal;
    }

    items(items: cog.Builder<dashboardv2beta1.GridLayoutItemKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }

    item(item: cog.Builder<dashboardv2beta1.GridLayoutItemKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        const itemResource = item.build();
        this.internal.spec.items.push(itemResource);
        return this;
    }
}

