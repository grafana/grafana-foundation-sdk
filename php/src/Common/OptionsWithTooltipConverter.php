<?php

namespace Grafana\Foundation\Common;

final class OptionsWithTooltipConverter
{
    public static function convert(\Grafana\Foundation\Common\OptionsWithTooltip $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\OptionsWithTooltipBuilder())',
        ];
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

