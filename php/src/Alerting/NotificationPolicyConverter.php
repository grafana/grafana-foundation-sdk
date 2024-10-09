<?php

namespace Grafana\Foundation\Alerting;

final class NotificationPolicyConverter
{
    public static function convert(\Grafana\Foundation\Alerting\NotificationPolicy $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\NotificationPolicyBuilder())',
        ];
            if ($input->continue !== null) {
    
        
    $buffer = 'continue(';
        $arg0 =\var_export($input->continue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupBy !== null && count($input->groupBy) >= 1) {
    
        
    $buffer = 'groupBy(';
        $tmparg0 = [];
        foreach ($input->groupBy as $arg1) {
        $tmpgroup_byarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpgroup_byarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupInterval !== null && $input->groupInterval !== "") {
    
        
    $buffer = 'groupInterval(';
        $arg0 =\var_export($input->groupInterval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupWait !== null && $input->groupWait !== "") {
    
        
    $buffer = 'groupWait(';
        $arg0 =\var_export($input->groupWait, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->match !== null) {
    
        
    $buffer = 'match(';
        $arg0 = "[";
        foreach ($input->match as $key => $arg1) {
            $tmpmatcharg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmpmatcharg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'matchRe(';
        $arg0 = "[";
        foreach ($input->matchRe as $key => $arg1) {
            $tmpmatch_rearg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmpmatch_rearg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->matchers) >= 1) {
    
        
    $buffer = 'matchers(';
        $tmparg0 = [];
        foreach ($input->matchers as $arg1) {
        $tmpmatchersarg1 = \Grafana\Foundation\Alerting\MatcherConverter::convert($arg1);
        $tmparg0[] = $tmpmatchersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->muteTimeIntervals !== null && count($input->muteTimeIntervals) >= 1) {
    
        
    $buffer = 'muteTimeIntervals(';
        $tmparg0 = [];
        foreach ($input->muteTimeIntervals as $arg1) {
        $tmpmute_time_intervalsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpmute_time_intervalsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->objectMatchers) >= 1) {
    
        
    $buffer = 'objectMatchers(';
        $tmparg0 = [];
        foreach ($input->objectMatchers as $arg1) {
        $tmptmpobject_matchersarg1 = [];
        foreach ($arg1 as $arg1Value) {
        $tmparg1arg1Value =\var_export($arg1Value, true);
        $tmptmpobject_matchersarg1[] = $tmparg1arg1Value;
        }
        $tmpobject_matchersarg1 = "[" . implode(", \n", $tmptmpobject_matchersarg1) . "]";
        $tmparg0[] = $tmpobject_matchersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->provenance !== "") {
    
        
    $buffer = 'provenance(';
        $arg0 =\var_export($input->provenance, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->receiver !== null && $input->receiver !== "") {
    
        
    $buffer = 'receiver(';
        $arg0 =\var_export($input->receiver, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeatInterval !== null && $input->repeatInterval !== "") {
    
        
    $buffer = 'repeatInterval(';
        $arg0 =\var_export($input->repeatInterval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->routes !== null && count($input->routes) >= 1) {
    
        
    $buffer = 'routes(';
        $tmparg0 = [];
        foreach ($input->routes as $arg1) {
        $tmproutesarg1 = \Grafana\Foundation\Alerting\NotificationPolicyConverter::convert($arg1);
        $tmparg0[] = $tmproutesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

