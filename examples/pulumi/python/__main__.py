from grafana_foundation_sdk.cog.encoder import JSONEncoder
import pulumiverse_grafana as grafana
import pulumi
import os
import sys


if __name__ == '__main__':
    # Hack to re-use other examples
    sys.path.append(os.path.join(os.path.dirname(__file__),
                    "..", '..', "python", "linux-node-overview"))
    from main import builder

    provider = grafana.Provider(
        'grafana', url='http://localhost:3000', auth='admin:admin')

    dashboard = builder.build()
    dashboard_json = JSONEncoder(
        sort_keys=True, indent=2).encode(dashboard)

    pulumi_dashboard = grafana.dashboard.Dashboard(dashboard.uid,
                                                   config_json=dashboard_json,
                                                   opts=pulumi.ResourceOptions(providers=[provider]))

    # Export the Name of the repository
    pulumi.export('dashboardUID', pulumi_dashboard.uid)
