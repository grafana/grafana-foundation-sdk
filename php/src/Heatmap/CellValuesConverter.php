<?php

namespace Grafana\Foundation\Heatmap;

final class CellValuesConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\CellValues $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\CellValuesBuilder())',
        ];
            if ($input->unit !== null && $input->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->decimals !== null) {
    
        
    $buffer = 'decimals(';
        $arg0 =\var_export($input->decimals, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

