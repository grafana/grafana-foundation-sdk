// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class GridBuilder implements cog.Builder<dashboardv2.GridLayoutKind> {
    protected readonly internal: dashboardv2.GridLayoutKind;

    constructor() {
        this.internal = dashboardv2.defaultGridLayoutKind();
        this.internal.kind = "GridLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.GridLayoutKind {
        return this.internal;
    }

    items(items: cog.Builder<dashboardv2.GridLayoutItemKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }

    item(item: cog.Builder<dashboardv2.GridLayoutItemKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        const itemResource = item.build();
        this.internal.spec.items.push(itemResource);
        return this;
    }
}

export function grid(): GridBuilder {
	const builder = new GridBuilder();

	return builder;
}

