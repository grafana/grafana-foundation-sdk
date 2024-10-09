<?php

namespace Grafana\Foundation\Heatmap;

final class HeatmapTooltipConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\HeatmapTooltip $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\HeatmapTooltipBuilder())',
        ];
            
    
        {
    $buffer = 'show(';
        $arg0 =\var_export($input->show, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->yHistogram !== null) {
    
        
    $buffer = 'yHistogram(';
        $arg0 =\var_export($input->yHistogram, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showColorScale !== null) {
    
        
    $buffer = 'showColorScale(';
        $arg0 =\var_export($input->showColorScale, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

