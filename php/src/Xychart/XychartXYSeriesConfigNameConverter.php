<?php

namespace Grafana\Foundation\Xychart;

final class XychartXYSeriesConfigNameConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XychartXYSeriesConfigName $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XychartXYSeriesConfigNameBuilder())',
        ];
            if ($input->fixed !== null && $input->fixed !== "") {
    
        
    $buffer = 'fixed(';
        $arg0 =\var_export($input->fixed, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

