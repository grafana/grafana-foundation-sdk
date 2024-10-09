<?php

namespace Grafana\Foundation\Heatmap;

final class RowsHeatmapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\RowsHeatmapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\RowsHeatmapOptionsBuilder())',
        ];
            if ($input->value !== null && $input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->layout !== null) {
    
        
    $buffer = 'layout(';
        $arg0 ='\Grafana\Foundation\Common\HeatmapCellLayout::fromValue("'.$input->layout.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

