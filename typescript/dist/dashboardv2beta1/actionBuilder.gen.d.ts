import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ActionBuilder implements cog.Builder<dashboardv2beta1.Action> {
    protected readonly internal: dashboardv2beta1.Action;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.Action;
    type(type: dashboardv2beta1.ActionType): this;
    title(title: string): this;
    fetch(fetch: cog.Builder<dashboardv2beta1.FetchOptions>): this;
    infinity(infinity: cog.Builder<dashboardv2beta1.InfinityOptions>): this;
    confirmation(confirmation: string): this;
    oneClick(oneClick: boolean): this;
    variables(variables: cog.Builder<dashboardv2beta1.ActionVariable>[]): this;
    style(style: {
        backgroundColor?: string;
    }): this;
}
