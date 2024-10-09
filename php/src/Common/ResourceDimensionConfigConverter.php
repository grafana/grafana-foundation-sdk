<?php

namespace Grafana\Foundation\Common;

final class ResourceDimensionConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\ResourceDimensionConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\ResourceDimensionConfigBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\ResourceDimensionMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
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

