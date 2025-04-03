<?php

namespace Grafana\Foundation\Common;

final class ScaleDistributionConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\ScaleDistributionConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\ScaleDistributionConfigBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Common\ScaleDistribution::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->log !== null) {
    
        
    $buffer = 'log(';
        $arg0 =\var_export($input->log, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->linearThreshold !== null) {
    
        
    $buffer = 'linearThreshold(';
        $arg0 =\var_export($input->linearThreshold, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

