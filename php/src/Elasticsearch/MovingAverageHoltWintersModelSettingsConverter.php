<?php

namespace Grafana\Foundation\Elasticsearch;

final class MovingAverageHoltWintersModelSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\MovingAverageHoltWintersModelSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\MovingAverageHoltWintersModelSettingsBuilder())',
        ];
            
    
        {
    $buffer = 'settings(';
        $arg0 = \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettingsConverter::convert($input->settings);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->window !== "") {
    
        
    $buffer = 'window(';
        $arg0 =\var_export($input->window, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'minimize(';
        $arg0 =\var_export($input->minimize, true);
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

