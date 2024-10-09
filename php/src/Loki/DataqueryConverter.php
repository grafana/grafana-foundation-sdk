<?php

namespace Grafana\Foundation\Loki;

final class DataqueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Loki\Dataquery);
        $calls = [
            '(new \Grafana\Foundation\Loki\DataqueryBuilder())',
        ];
            if ($input->expr !== "") {
    
        
    $buffer = 'expr(';
        $arg0 =\var_export($input->expr, true);
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
            if ($input->maxLines !== null) {
    
        
    $buffer = 'maxLines(';
        $arg0 =\var_export($input->maxLines, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resolution !== null) {
    
        
    $buffer = 'resolution(';
        $arg0 =\var_export($input->resolution, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->editorMode !== null) {
    
        
    $buffer = 'editorMode(';
        $arg0 ='\Grafana\Foundation\Loki\QueryEditorMode::fromValue("'.$input->editorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->range !== null) {
    
        
    $buffer = 'range(';
        $arg0 =\var_export($input->range, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->instant !== null) {
    
        
    $buffer = 'instant(';
        $arg0 =\var_export($input->instant, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->step !== null && $input->step !== "") {
    
        
    $buffer = 'step(';
        $arg0 =\var_export($input->step, true);
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

        return \implode("\n\t->", $calls);
    }
}

