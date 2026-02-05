import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class PanelBuilder implements cog.Builder<dashboardv2beta1.PanelKind> {
    protected readonly internal: dashboardv2beta1.PanelKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.PanelKind;
    id(id: number): this;
    title(title: string): this;
    description(description: string): this;
    links(links: cog.Builder<dashboardv2beta1.DataLink>[]): this;
    data(data: cog.Builder<dashboardv2beta1.QueryGroupKind>): this;
    visualization(vizConfig: cog.Builder<dashboardv2beta1.VizConfigKind>): this;
    transparent(transparent: boolean): this;
}
