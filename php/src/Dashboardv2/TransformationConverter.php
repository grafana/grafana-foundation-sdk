<?php

namespace Grafana\Foundation\Dashboardv2;

final class TransformationConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\TransformationKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\TransformationBuilder())',
        ];
            if ($input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->disabled !== null) {
    
        
    $buffer = 'disabled(';
        $arg0 =\var_export($input->spec->disabled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->filter !== null) {
    
        
    $buffer = 'filter(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2\MatcherConfig(id: '.\var_export($input->spec->filter->id, true).','.(($input->spec->filter->scope !== null) ? 'scope: '.'\Grafana\Foundation\Dashboardv2\MatcherScope::fromValue("'.$input->spec->filter->scope.'")'.', ' : '').''.(($input->spec->filter->options !== null) ? 'options: '.\var_export($input->spec->filter->options, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->topic !== null) {
    
        
    $buffer = 'topic(';
        $arg0 ='\Grafana\Foundation\Dashboardv2\DataTopic::fromValue("'.$input->spec->topic.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null) {
    
        
    $buffer = 'options(';
        $arg0 =\var_export($input->spec->options, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

