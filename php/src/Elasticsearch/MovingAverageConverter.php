<?php

namespace Grafana\Foundation\Elasticsearch;

final class MovingAverageConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\MovingAverage $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\MovingAverageBuilder())',
        ];
            if ($input->pipelineAgg !== null && $input->pipelineAgg !== "") {
    
        
    $buffer = 'pipelineAgg(';
        $arg0 =\var_export($input->pipelineAgg, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->settings !== null) {
    
        
    $buffer = 'settings(';
        $arg0 = "[";
        foreach ($input->settings as $key => $arg1) {
            $tmpsettingsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmpsettingsarg1,";
        }
        $arg0 .= "]";
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

        return \implode("\n\t->", $calls);
    }
}

