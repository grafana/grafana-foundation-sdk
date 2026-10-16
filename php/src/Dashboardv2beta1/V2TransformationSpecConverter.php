<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class V2TransformationSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\V2TransformationSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\V2TransformationSpecBuilder())',
        ];
            if ($input->disabled !== null) {
    
        
    $buffer = 'disabled(';
        $arg0 =\var_export($input->disabled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filter !== null) {
    
        
    $buffer = 'filter(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(id: '.\var_export($input->filter->id, true).','.(($input->filter->scope !== null) ? 'scope: '.'\Grafana\Foundation\Dashboardv2beta1\MatcherScope::fromValue("'.$input->filter->scope.'")'.', ' : '').''.(($input->filter->options !== null) ? 'options: '.\var_export($input->filter->options, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->topic !== null) {
    
        
    $buffer = 'topic(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\DataTopic::fromValue("'.$input->topic.'")';
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

