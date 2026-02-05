<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TransformationConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TransformationKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TransformationBuilder())',
        ];
            if ($input->kind !== "") {
    
        
    $buffer = 'kind(';
        $arg0 =\var_export($input->kind, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->spec->id, true);
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
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(id: '.\var_export($input->spec->filter->id, true).','.(($input->spec->filter->options !== null) ? 'options: '.\var_export($input->spec->filter->options, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->topic !== null) {
    
        
    $buffer = 'topic(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\DataTopic::fromValue("'.$input->spec->topic.'")';
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

