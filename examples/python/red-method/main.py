from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.models.resource import Manifest, Metadata, DashboardKind

from src.red import red


red_dashboard = red(
    dashboard_title="RED method",
    service_ids=["sample-service", "payments", "front-gateway"],
)

if __name__ == "__main__":
    dashboard = red_dashboard.build()

    manifest = Manifest(
        api_version="dashboard.grafana.app/v1beta1",
        kind=DashboardKind,
        metadata=Metadata(name=dashboard.uid),
        spec=dashboard,
    )

    print(JSONEncoder(sort_keys=True, indent=2).encode(manifest))
