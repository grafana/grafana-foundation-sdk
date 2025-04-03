<?php

namespace Grafana\Foundation\Xychart;

final class XychartXYSeriesConfigXConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XychartXYSeriesConfigX $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XychartXYSeriesConfigXBuilder())',
        ];
            
    
        {
    $buffer = 'matcher(';
        $arg0 = \Grafana\Foundation\Xychart\MatcherConfigConverter::convert($input->matcher);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

