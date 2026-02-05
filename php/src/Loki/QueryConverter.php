<?php

namespace Grafana\Foundation\Loki;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Loki\QueryBuilder())',
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->expr !== "") {
    
        
    $buffer = 'expr(';
        $arg0 =\var_export($input->spec->expr, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->legendFormat !== null && $input->spec->legendFormat !== "") {
    
        
    $buffer = 'legendFormat(';
        $arg0 =\var_export($input->spec->legendFormat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->maxLines !== null) {
    
        
    $buffer = 'maxLines(';
        $arg0 =\var_export($input->spec->maxLines, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->resolution !== null) {
    
        
    $buffer = 'resolution(';
        $arg0 =\var_export($input->spec->resolution, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->editorMode !== null) {
    
        
    $buffer = 'editorMode(';
        $arg0 ='\Grafana\Foundation\Loki\QueryEditorMode::fromValue("'.$input->spec->editorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->range !== null) {
    
        
    $buffer = 'range(';
        $arg0 =\var_export($input->spec->range, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->instant !== null) {
    
        
    $buffer = 'instant(';
        $arg0 =\var_export($input->spec->instant, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->step !== null && $input->spec->step !== "") {
    
        
    $buffer = 'step(';
        $arg0 =\var_export($input->spec->step, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Loki\Dataquery && $input->spec->direction !== null) {
    
        
    $buffer = 'direction(';
        $arg0 ='\Grafana\Foundation\Loki\LokiQueryDirection::fromValue("'.$input->spec->direction.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

