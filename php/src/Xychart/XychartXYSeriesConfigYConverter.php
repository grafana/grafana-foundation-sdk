<?php

namespace Grafana\Foundation\Xychart;

final class XychartXYSeriesConfigYConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XychartXYSeriesConfigY $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XychartXYSeriesConfigYBuilder())',
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

