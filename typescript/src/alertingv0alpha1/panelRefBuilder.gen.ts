// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class PanelRefBuilder implements cog.Builder<alertingv0alpha1.PanelRef> {
    protected readonly internal: alertingv0alpha1.PanelRef;

    constructor() {
        this.internal = alertingv0alpha1.defaultPanelRef();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.PanelRef {
        return this.internal;
    }

    dashboardUID(dashboardUID: string): this {
        if (!(dashboardUID.length >= 1)) {
            throw new Error("dashboardUID.length must be >= 1");
        }
        if (!(new RegExp("^[a-zA-Z0-9_-]+$").test(dashboardUID))) {
            throw new Error("dashboardUID must match regex ^[a-zA-Z0-9_-]+$");
        }
        this.internal.dashboardUID = dashboardUID;
        return this;
    }

    panelID(panelID: number): this {
        if (!(panelID > 0)) {
            throw new Error("panelID must be > 0");
        }
        this.internal.panelID = panelID;
        return this;
    }
}

