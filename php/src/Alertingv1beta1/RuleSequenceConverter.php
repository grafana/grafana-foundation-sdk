<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class RuleSequenceConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\RuleSequenceSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\RuleSequenceBuilder())',
        ];
            
    
        {
    $buffer = 'trigger(';
        $arg0 = \Grafana\Foundation\Alertingv1beta1\IntervalTriggerConverter::convert($input->trigger);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->recordingRules) >= 1) {
    
        
    $buffer = 'recordingRules(';
        $tmparg0 = [];
        foreach ($input->recordingRules as $arg1) {
        $tmprecordingRulesarg1 = \Grafana\Foundation\Alertingv1beta1\RuleRefConverter::convert($arg1);
        $tmparg0[] = $tmprecordingRulesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alertRules !== null && count($input->alertRules) >= 1) {
    
        
    $buffer = 'alertRules(';
        $tmparg0 = [];
        foreach ($input->alertRules as $arg1) {
        $tmpalertRulesarg1 = \Grafana\Foundation\Alertingv1beta1\RuleRefConverter::convert($arg1);
        $tmparg0[] = $tmpalertRulesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

