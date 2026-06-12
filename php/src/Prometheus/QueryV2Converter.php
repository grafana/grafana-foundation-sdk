<?php

namespace Grafana\Foundation\Prometheus;

final class QueryV2Converter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Prometheus\QueryV2Builder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->labels !== null) {
    
        
    $buffer = 'labels(';
        $arg0 = "[";
        foreach ($input->labels as $key => $arg1) {
            $tmplabelsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmplabelsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->adhocFilters !== null && count($input->spec->adhocFilters) >= 1) {
    
        
    $buffer = 'adhocFilters(';
        $tmparg0 = [];
        foreach ($input->spec->adhocFilters as $arg1) {
        $tmpadhocFiltersarg1 = \Grafana\Foundation\Prometheus\AdhocFiltersConverter::convert($arg1);
        $tmparg0[] = $tmpadhocFiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->editorMode !== null) {
    
        
    $buffer = 'editorMode(';
        $arg0 ='\Grafana\Foundation\Prometheus\QueryEditorMode::fromValue("'.$input->spec->editorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->exemplar !== null) {
    
        
    $buffer = 'exemplar(';
        $arg0 =\var_export($input->spec->exemplar, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->expr !== "") {
    
        
    $buffer = 'expr(';
        $arg0 =\var_export($input->spec->expr, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->format !== null) {
    
        
    $buffer = 'format(';
        $arg0 ='\Grafana\Foundation\Prometheus\PromQueryFormat::fromValue("'.$input->spec->format.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->groupByKeys !== null && count($input->spec->groupByKeys) >= 1) {
    
        
    $buffer = 'groupByKeys(';
        $tmparg0 = [];
        foreach ($input->spec->groupByKeys as $arg1) {
        $tmpgroupByKeysarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpgroupByKeysarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->instant !== null) {
    
        
    $buffer = 'instant(';
        $arg0 =\var_export($input->spec->instant, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->intervalFactor !== null) {
    
        
    $buffer = 'intervalFactor(';
        $arg0 =\var_export($input->spec->intervalFactor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->intervalMs !== null) {
    
        
    $buffer = 'intervalMs(';
        $arg0 =\var_export($input->spec->intervalMs, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->legendFormat !== null && $input->spec->legendFormat !== "") {
    
        
    $buffer = 'legendFormat(';
        $arg0 =\var_export($input->spec->legendFormat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->maxDataPoints !== null) {
    
        
    $buffer = 'maxDataPoints(';
        $arg0 =\var_export($input->spec->maxDataPoints, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->range !== null) {
    
        
    $buffer = 'range(';
        $arg0 =\var_export($input->spec->range, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->resultAssertions !== null) {
    
        
    $buffer = 'resultAssertions(';
        $arg0 = \Grafana\Foundation\Prometheus\ResultAssertionsConverter::convert($input->spec->resultAssertions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->scopes !== null && count($input->spec->scopes) >= 1) {
    
        
    $buffer = 'scopes(';
        $tmparg0 = [];
        foreach ($input->spec->scopes as $arg1) {
        $tmpscopesarg1 = \Grafana\Foundation\Prometheus\ScopesConverter::convert($arg1);
        $tmparg0[] = $tmpscopesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->timeRange !== null) {
    
        
    $buffer = 'timeRange(';
        $arg0 = \Grafana\Foundation\Prometheus\TimeRangeConverter::convert($input->spec->timeRange);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->interval !== null && $input->spec->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->spec->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

