<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class RuleSequenceConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\RuleSequenceSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\RuleSequenceBuilder())',
        ];
            
    
        {
    $buffer = 'trigger(';
        $arg0 = \Grafana\Foundation\Alertingv0alpha1\IntervalTriggerConverter::convert($input->trigger);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->recordingRules) >= 1) {
    
        
    $buffer = 'recordingRules(';
        $tmparg0 = [];
        foreach ($input->recordingRules as $arg1) {
        $tmprecordingRulesarg1 = \Grafana\Foundation\Alertingv0alpha1\RuleRefConverter::convert($arg1);
        $tmparg0[] = $tmprecordingRulesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alertingRules !== null && count($input->alertingRules) >= 1) {
    
        
    $buffer = 'alertingRules(';
        $tmparg0 = [];
        foreach ($input->alertingRules as $arg1) {
        $tmpalertingRulesarg1 = \Grafana\Foundation\Alertingv0alpha1\RuleRefConverter::convert($arg1);
        $tmparg0[] = $tmpalertingRulesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

