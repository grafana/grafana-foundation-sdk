<?php

namespace Grafana\Foundation\Prometheus;

final class DataqueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $calls = [
            '(new \Grafana\Foundation\Prometheus\DataqueryBuilder())',
        ];
            if ($input->expr !== "") {
    
        
    $buffer = 'expr(';
        $arg0 =\var_export($input->expr, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->instant !== null && $input->instant === true && $input->range !== null && $input->range === false) {
    
        
    $buffer = 'instant(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->range !== null && $input->range === true && $input->instant !== null && $input->instant === false) {
    
        
    $buffer = 'range(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->exemplar !== null) {
    
        
    $buffer = 'exemplar(';
        $arg0 =\var_export($input->exemplar, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->editorMode !== null) {
    
        
    $buffer = 'editorMode(';
        $arg0 ='\Grafana\Foundation\Prometheus\QueryEditorMode::fromValue("'.$input->editorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->format !== null) {
    
        
    $buffer = 'format(';
        $arg0 ='\Grafana\Foundation\Prometheus\PromQueryFormat::fromValue("'.$input->format.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->legendFormat !== null && $input->legendFormat !== "") {
    
        
    $buffer = 'legendFormat(';
        $arg0 =\var_export($input->legendFormat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->intervalFactor !== null) {
    
        
    $buffer = 'intervalFactor(';
        $arg0 =\var_export($input->intervalFactor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->interval !== null && $input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

