<?php

namespace Grafana\Foundation\Heatmap;

final class HeatmapColorOptionsConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\HeatmapColorOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\HeatmapColorOptionsBuilder())',
        ];
            if ($input->mode !== null) {
    
        
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Heatmap\HeatmapColorMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->scheme !== "") {
    
        
    $buffer = 'scheme(';
        $arg0 =\var_export($input->scheme, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fill !== "") {
    
        
    $buffer = 'fill(';
        $arg0 =\var_export($input->fill, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->scale !== null) {
    
        
    $buffer = 'scale(';
        $arg0 ='\Grafana\Foundation\Heatmap\HeatmapColorScale::fromValue("'.$input->scale.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'exponent(';
        $arg0 =\var_export($input->exponent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'steps(';
        $arg0 =\var_export($input->steps, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'reverse(';
        $arg0 =\var_export($input->reverse, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

