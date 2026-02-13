import { red } from "./red";

const redDashboard = red({
  dashboardTitle: "RED method",
  serviceIds: ["sample-service", "payments", "front-gateway"],
});

const dashboard = redDashboard.build();

console.log(JSON.stringify({
    apiVersion: 'dashboard.grafana.app/v1beta1',
    kind: 'Dashboard',
    metadata: {
        name: dashboard.uid!,
    },
    spec: dashboard,
}, null, 2));
