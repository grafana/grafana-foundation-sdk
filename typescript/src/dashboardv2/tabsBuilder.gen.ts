// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class TabsBuilder implements cog.Builder<dashboardv2.TabsLayoutKind> {
    protected readonly internal: dashboardv2.TabsLayoutKind;

    constructor() {
        this.internal = dashboardv2.defaultTabsLayoutKind();
        this.internal.kind = "TabsLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.TabsLayoutKind {
        return this.internal;
    }

    tabs(tabs: cog.Builder<dashboardv2.TabsLayoutTabKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTabsLayoutSpec();
        }
        const tabsResources = tabs.map(builder1 => builder1.build());
        this.internal.spec.tabs = tabsResources;
        return this;
    }

    tab(tab: cog.Builder<dashboardv2.TabsLayoutTabKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTabsLayoutSpec();
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

