---
title: <span class="badge object-type-enum"></span> TestDataQueryType
---
# <span class="badge object-type-enum"></span> TestDataQueryType

## Definition

```python
class TestDataQueryType(enum.StrEnum):
    RANDOM_WALK = "random_walk"
    SLOW_QUERY = "slow_query"
    RANDOM_WALK_WITH_ERROR = "random_walk_with_error"
    RANDOM_WALK_TABLE = "random_walk_table"
    EXPONENTIAL_HEATMAP_BUCKET_DATA = "exponential_heatmap_bucket_data"
    LINEAR_HEATMAP_BUCKET_DATA = "linear_heatmap_bucket_data"
    NO_DATA_POINTS = "no_data_points"
    DATA_POINTS_OUTSIDE_RANGE = "datapoints_outside_range"
    CSV_METRIC_VALUES = "csv_metric_values"
    PREDICTABLE_PULSE = "predictable_pulse"
    PREDICTABLE_CSV_WAVE = "predictable_csv_wave"
    STREAMING_CLIENT = "streaming_client"
    SIMULATION = "simulation"
    USA = "usa"
    LIVE = "live"
    GRAFANA_API = "grafana_api"
    ARROW = "arrow"
    ANNOTATIONS = "annotations"
    TABLE_STATIC = "table_static"
    SERVER_ERROR500 = "server_error_500"
    LOGS = "logs"
    NODE_GRAPH = "node_graph"
    FLAME_GRAPH = "flame_graph"
    RAW_FRAME = "raw_frame"
    CSV_FILE = "csv_file"
    CSV_CONTENT = "csv_content"
    TRACE = "trace"
    MANUAL_ENTRY = "manual_entry"
    VARIABLES_QUERY = "variables-query"
```
