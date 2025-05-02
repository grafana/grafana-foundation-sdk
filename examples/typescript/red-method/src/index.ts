import { red } from "./red";

const redDashboard = red({
  dashboardTitle: "RED method",
  serviceIds: ["sample-service", "payments", "front-gateway"],
});

console.log(JSON.stringify(redDashboard.build(), null, 2));
