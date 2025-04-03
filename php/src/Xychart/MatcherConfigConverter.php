<?php

namespace Grafana\Foundation\Xychart;

final class MatcherConfigConverter
{
    public static function convert(\Grafana\Foundation\Xychart\MatcherConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\MatcherConfigBuilder())',
        ];
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null) {
    
        
    $buffer = 'options(';
        $arg0 =\var_export($input->options, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

