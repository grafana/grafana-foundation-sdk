from grafana_foundation_sdk.cog.encoder import JSONEncoder

from src.red import red


if __name__ == '__main__':
    red_dashboard = red(
        dashboard_title="RED method",
        service_ids=["sample-service", "payments", "front-gateway"],
    )

    print(JSONEncoder(sort_keys=True, indent=2).encode(red_dashboard.build()))
