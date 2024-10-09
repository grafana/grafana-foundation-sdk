<?php

namespace Grafana\Foundation\Common;

final class HideSeriesConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\HideSeriesConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\HideSeriesConfigBuilder())',
        ];
            
    
        {
    $buffer = 'tooltip(';
        $arg0 =\var_export($input->tooltip, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'legend(';
        $arg0 =\var_export($input->legend, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'viz(';
        $arg0 =\var_export($input->viz, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

