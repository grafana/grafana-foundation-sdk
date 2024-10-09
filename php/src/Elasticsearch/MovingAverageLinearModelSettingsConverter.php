<?php

namespace Grafana\Foundation\Elasticsearch;

final class MovingAverageLinearModelSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\MovingAverageLinearModelSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\MovingAverageLinearModelSettingsBuilder())',
        ];
            if ($input->window !== "") {
    
        
    $buffer = 'window(';
        $arg0 =\var_export($input->window, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->predict !== "") {
    
        
    $buffer = 'predict(';
        $arg0 =\var_export($input->predict, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

