import * as pulumi from "@pulumi/pulumi";
import * as grafana from "@pulumiverse/grafana";
import { builder } from '../../typescript/linux-node-overview/src/index';

const provider = new grafana.Provider("grafana", { url: "http://localhost:3000", auth: "admin:admin" });

const dashboard = builder.build();

const pulumiDashboard = new grafana.Dashboard(dashboard.uid || 'uid', { configJson: JSON.stringify(dashboard, null, 2) }, { provider: provider });

export const dashboardUid = pulumiDashboard.uid;
