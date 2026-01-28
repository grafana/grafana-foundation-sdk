// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TabsLayoutSpecBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutSpec> {
    protected readonly internal: dashboardv2beta1.TabsLayoutSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultTabsLayoutSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutSpec {
        return this.internal;
    }

    tabs(tabs: cog.Builder<dashboardv2beta1.TabsLayoutTabKind>[]): this {
        const tabsResources = tabs.map(builder1 => builder1.build());
        this.internal.tabs = tabsResources;
        return this;
    }
}

