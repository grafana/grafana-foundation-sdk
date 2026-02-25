// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TabsBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutKind> {
    protected readonly internal: dashboardv2beta1.TabsLayoutKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultTabsLayoutKind();
        this.internal.kind = "TabsLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutKind {
        return this.internal;
    }

    tabs(tabs: cog.Builder<dashboardv2beta1.TabsLayoutTabKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutSpec();
        }
        const tabsResources = tabs.map(builder1 => builder1.build());
        this.internal.spec.tabs = tabsResources;
        return this;
    }

    tab(tab: cog.Builder<dashboardv2beta1.TabsLayoutTabKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutSpec();
        }
        if (!this.internal.spec.tabs) {
            this.internal.spec.tabs = [];
        }
        const tabResource = tab.build();
        this.internal.spec.tabs.push(tabResource);
        return this;
    }
}

export function tabs(): TabsBuilder {
	const builder = new TabsBuilder();

	return builder;
}

