import * as cog from '../cog';
import * as dashboard from '../dashboard';
import * as common from '../common';
export declare class RowBuilder implements cog.Builder<dashboard.RowPanel> {
    protected readonly internal: dashboard.RowPanel;
    constructor(title: string);
    /**
     * Builds the object.
     */
    build(): dashboard.RowPanel;
    collapsed(collapsed: boolean): this;
    title(title: string): this;
    datasource(datasource: common.DataSourceRef): this;
    gridPos(gridPos: dashboard.GridPos): this;
    id(id: number): this;
    withPanel(panel: cog.Builder<dashboard.Panel>): this;
    repeat(repeat: string): this;
}
