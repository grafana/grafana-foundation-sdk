<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class NamedRoutingTreeConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\NamedRoutingTree $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\NamedRoutingTreeBuilder())',
        ];
            if ($input->routingTree !== "") {
    
        
    $buffer = 'routingTree(';
        $arg0 =\var_export($input->routingTree, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

