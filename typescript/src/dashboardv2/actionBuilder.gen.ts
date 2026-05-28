// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ActionBuilder implements cog.Builder<dashboardv2.Action> {
    protected readonly internal: dashboardv2.Action;

    constructor() {
        this.internal = dashboardv2.defaultAction();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.Action {
        return this.internal;
    }

    type(type: dashboardv2.ActionType): this {
        this.internal.type = type;
        return this;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    fetch(fetch: cog.Builder<dashboardv2.FetchOptions>): this {
        const fetchResource = fetch.build();
        this.internal.fetch = fetchResource;
        return this;
    }

    infinity(infinity: cog.Builder<dashboardv2.InfinityOptions>): this {
        const infinityResource = infinity.build();
        this.internal.infinity = infinityResource;
        return this;
    }

    confirmation(confirmation: string): this {
        this.internal.confirmation = confirmation;
        return this;
    }

    oneClick(oneClick: boolean): this {
        this.internal.oneClick = oneClick;
        return this;
    }

    variables(variables: cog.Builder<dashboardv2.ActionVariable>[]): this {
        const variablesResources = variables.map(builder1 => builder1.build());
        this.internal.variables = variablesResources;
        return this;
    }

    style(style: {
	backgroundColor?: string;
}): this {
        this.internal.style = style;
        return this;
    }
}

