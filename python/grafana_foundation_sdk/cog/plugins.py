# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import alertgroups
from ..models import annotationslist
from ..models import azuremonitor
from ..models import barchart
from ..models import bargauge
from ..models import candlestick
from ..models import canvas
from ..models import cloudwatch
from ..models import dashboardlist
from ..models import datagrid
from ..models import debug
from ..models import elasticsearch
from ..models import expr
from ..models import gauge
from ..models import geomap
from ..models import googlecloudmonitoring
from ..models import grafanapyroscope
from ..models import heatmap
from ..models import histogram
from ..models import logs
from ..models import loki
from ..models import news
from ..models import nodegraph
from ..models import parca
from ..models import piechart
from ..models import prometheus
from ..models import stat
from ..models import statetimeline
from ..models import statushistory
from ..models import table
from ..models import tempo
from ..models import testdata
from ..models import text
from ..models import timeseries
from ..models import trend
from ..models import xychart
from . import runtime as cogruntime


def register_default_plugins():
    # Panelcfg variants
    cogruntime.register_panelcfg_variant(alertgroups.variant_config())
    cogruntime.register_panelcfg_variant(annotationslist.variant_config())
    cogruntime.register_panelcfg_variant(barchart.variant_config())
    cogruntime.register_panelcfg_variant(bargauge.variant_config())
    cogruntime.register_panelcfg_variant(candlestick.variant_config())
    cogruntime.register_panelcfg_variant(canvas.variant_config())
    cogruntime.register_panelcfg_variant(dashboardlist.variant_config())
    cogruntime.register_panelcfg_variant(datagrid.variant_config())
    cogruntime.register_panelcfg_variant(debug.variant_config())
    cogruntime.register_panelcfg_variant(gauge.variant_config())
    cogruntime.register_panelcfg_variant(geomap.variant_config())
    cogruntime.register_panelcfg_variant(heatmap.variant_config())
    cogruntime.register_panelcfg_variant(histogram.variant_config())
    cogruntime.register_panelcfg_variant(logs.variant_config())
    cogruntime.register_panelcfg_variant(news.variant_config())
    cogruntime.register_panelcfg_variant(nodegraph.variant_config())
    cogruntime.register_panelcfg_variant(piechart.variant_config())
    cogruntime.register_panelcfg_variant(stat.variant_config())
    cogruntime.register_panelcfg_variant(statetimeline.variant_config())
    cogruntime.register_panelcfg_variant(statushistory.variant_config())
    cogruntime.register_panelcfg_variant(table.variant_config())
    cogruntime.register_panelcfg_variant(text.variant_config())
    cogruntime.register_panelcfg_variant(timeseries.variant_config())
    cogruntime.register_panelcfg_variant(trend.variant_config())
    cogruntime.register_panelcfg_variant(xychart.variant_config())

    # Dataquery variants
    cogruntime.register_dataquery_variant(azuremonitor.variant_config())
    cogruntime.register_dataquery_variant(cloudwatch.variant_config())
    cogruntime.register_dataquery_variant(elasticsearch.variant_config())
    cogruntime.register_dataquery_variant(expr.variant_config())
    cogruntime.register_dataquery_variant(googlecloudmonitoring.variant_config())
    cogruntime.register_dataquery_variant(grafanapyroscope.variant_config())
    cogruntime.register_dataquery_variant(loki.variant_config())
    cogruntime.register_dataquery_variant(parca.variant_config())
    cogruntime.register_dataquery_variant(prometheus.variant_config())
    cogruntime.register_dataquery_variant(tempo.variant_config())
    cogruntime.register_dataquery_variant(testdata.variant_config())
