<?php

namespace Grafana\Foundation\Xychart;

final class XYDimensionConfigConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XYDimensionConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XYDimensionConfigBuilder())',
        ];
            
    
        {
    $buffer = 'frame(';
        $arg0 =\var_export($input->frame, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->x !== null && $input->x !== "") {
    
        
    $buffer = 'x(';
        $arg0 =\var_export($input->x, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->exclude !== null && count($input->exclude) >= 1) {
    
        
    $buffer = 'exclude(';
        $tmparg0 = [];
        foreach ($input->exclude as $arg1) {
        $tmpexcludearg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpexcludearg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

