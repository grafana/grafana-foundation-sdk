<?php

namespace Grafana\Foundation\Prometheus;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Prometheus\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->instant !== null && $input->spec->instant === true && $input->spec->range !== null && $input->spec->range === false) {
    
        
    $buffer = 'instant(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->range !== null && $input->spec->range === true && $input->spec->instant !== null && $input->spec->instant === false) {
    
        
    $buffer = 'range(';
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->editorMode !== null) {
    
        
    $buffer = 'editorMode(';
        $arg0 ='\Grafana\Foundation\Prometheus\QueryEditorMode::fromValue("'.$input->spec->editorMode.'")';
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->legendFormat !== null && $input->spec->legendFormat !== "") {
    
        
    $buffer = 'legendFormat(';
        $arg0 =\var_export($input->spec->legendFormat, true);
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Prometheus\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
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

