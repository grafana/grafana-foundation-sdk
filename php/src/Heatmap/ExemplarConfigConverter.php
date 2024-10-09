<?php

namespace Grafana\Foundation\Heatmap;

final class ExemplarConfigConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\ExemplarConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\ExemplarConfigBuilder())',
        ];
            if ($input->color !== "") {
    
        
    $buffer = 'color(';
        $arg0 =\var_export($input->color, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

