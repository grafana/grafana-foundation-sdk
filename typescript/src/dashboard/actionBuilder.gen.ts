// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Dashboard action
export class ActionBuilder implements cog.Builder<dashboard.Action> {
    protected readonly internal: dashboard.Action;

    constructor() {
        this.internal = dashboard.defaultAction();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.Action {
        return this.internal;
    }

    type(type: dashboard.ActionType): this {
        this.internal.type = type;
        return this;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    fetch(fetch: cog.Builder<dashboard.FetchOptions>): this {
        const fetchResource = fetch.build();
        this.internal.fetch = fetchResource;
        return this;
    }

    infinity(infinity: cog.Builder<dashboard.InfinityOptions>): this {
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

    variables(variables: cog.Builder<dashboard.ActionVariable>[]): this {
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

