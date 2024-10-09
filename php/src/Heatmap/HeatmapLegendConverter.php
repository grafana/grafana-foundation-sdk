<?php

namespace Grafana\Foundation\Heatmap;

final class HeatmapLegendConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\HeatmapLegend $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\HeatmapLegendBuilder())',
        ];
            
    
        {
    $buffer = 'show(';
        $arg0 =\var_export($input->show, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

