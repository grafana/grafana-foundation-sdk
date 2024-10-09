<?php

namespace Grafana\Foundation\Common;

final class VizTooltipOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\VizTooltipOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\VizTooltipOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\TooltipDisplayMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'sort(';
        $arg0 ='\Grafana\Foundation\Common\SortOrder::fromValue("'.$input->sort.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

