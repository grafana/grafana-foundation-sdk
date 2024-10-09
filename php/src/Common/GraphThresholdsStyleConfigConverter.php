<?php

namespace Grafana\Foundation\Common;

final class GraphThresholdsStyleConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\GraphThresholdsStyleConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\GraphThresholdsStyleConfigBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\GraphTresholdsStyleMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

