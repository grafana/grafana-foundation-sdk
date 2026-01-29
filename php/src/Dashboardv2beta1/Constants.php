<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Constants
{
    /**
     * Annotation Query placement. Defines where the annotation query should be displayed.
     * - "inControlsMenu" renders the annotation query in the dashboard controls dropdown menu
     */
    const ANNOTATION_QUERY_PLACEMENT = "inControlsMenu";
    /**
     * Action variable type
     */
    const ACTION_VARIABLE_TYPE = "string";
    /**
     * Dashboard Link placement. Defines where the link should be displayed.
     * - "inControlsMenu" renders the link in bottom part of the dashboard controls dropdown menu
     */
    const DASHBOARD_LINK_PLACEMENT = "inControlsMenu";
    /**
     * Determine the origin of the adhoc variable filter
     */
    const FILTER_ORIGIN = "dashboard";
}