<?php

namespace Grafana\Foundation\Trend;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Trend\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Trend\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->xField !== null && $input->xField !== "") {
    
        
    $buffer = 'xField(';
        $arg0 =\var_export($input->xField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

