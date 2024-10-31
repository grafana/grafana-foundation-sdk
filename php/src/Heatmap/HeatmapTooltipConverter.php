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
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\TooltipDisplayMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->maxHeight !== null) {
    
        
    $buffer = 'maxHeight(';
        $arg0 =\var_export($input->maxHeight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxWidth !== null) {
    
        
    $buffer = 'maxWidth(';
        $arg0 =\var_export($input->maxWidth, true);
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

