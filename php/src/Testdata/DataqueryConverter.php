<?php

namespace Grafana\Foundation\Testdata;

final class DataqueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Testdata\Dataquery);
        $calls = [
            '(new \Grafana\Foundation\Testdata\DataqueryBuilder())',
        ];
            if ($input->alias !== null && $input->alias !== "") {
    
        
    $buffer = 'alias(';
        $arg0 =\var_export($input->alias, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->scenarioId !== null) {
    
        
    $buffer = 'scenarioId(';
        $arg0 ='\Grafana\Foundation\Testdata\TestDataQueryType::fromValue("'.$input->scenarioId.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->stringInput !== null && $input->stringInput !== "") {
    
        
    $buffer = 'stringInput(';
        $arg0 =\var_export($input->stringInput, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->stream !== null) {
    
        
    $buffer = 'stream(';
        $arg0 = \Grafana\Foundation\Testdata\StreamingQueryConverter::convert($input->stream);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pulseWave !== null) {
    
        
    $buffer = 'pulseWave(';
        $arg0 = \Grafana\Foundation\Testdata\PulseWaveQueryConverter::convert($input->pulseWave);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sim !== null) {
    
        
    $buffer = 'sim(';
        $arg0 = \Grafana\Foundation\Testdata\SimulationQueryConverter::convert($input->sim);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->csvWave !== null && count($input->csvWave) >= 1) {
    
        
    $buffer = 'csvWave(';
        $tmparg0 = [];
        foreach ($input->csvWave as $arg1) {
        $tmpcsvWavearg1 = \Grafana\Foundation\Testdata\CSVWaveConverter::convert($arg1);
        $tmparg0[] = $tmpcsvWavearg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->labels !== null && $input->labels !== "") {
    
        
    $buffer = 'labels(';
        $arg0 =\var_export($input->labels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lines !== null) {
    
        
    $buffer = 'lines(';
        $arg0 =\var_export($input->lines, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->levelColumn !== null) {
    
        
    $buffer = 'levelColumn(';
        $arg0 =\var_export($input->levelColumn, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->channel !== null && $input->channel !== "") {
    
        
    $buffer = 'channel(';
        $arg0 =\var_export($input->channel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->nodes !== null) {
    
        
    $buffer = 'nodes(';
        $arg0 = \Grafana\Foundation\Testdata\NodesQueryConverter::convert($input->nodes);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->csvFileName !== null && $input->csvFileName !== "") {
    
        
    $buffer = 'csvFileName(';
        $arg0 =\var_export($input->csvFileName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->csvContent !== null && $input->csvContent !== "") {
    
        
    $buffer = 'csvContent(';
        $arg0 =\var_export($input->csvContent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->rawFrameContent !== null && $input->rawFrameContent !== "") {
    
        
    $buffer = 'rawFrameContent(';
        $arg0 =\var_export($input->rawFrameContent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->seriesCount !== null) {
    
        
    $buffer = 'seriesCount(';
        $arg0 =\var_export($input->seriesCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->usa !== null) {
    
        
    $buffer = 'usa(';
        $arg0 = \Grafana\Foundation\Testdata\USAQueryConverter::convert($input->usa);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->errorType !== null) {
    
        
    $buffer = 'errorType(';
        $arg0 ='\Grafana\Foundation\Testdata\DataqueryErrorType::fromValue("'.$input->errorType.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spanCount !== null) {
    
        
    $buffer = 'spanCount(';
        $arg0 =\var_export($input->spanCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->points !== null && count($input->points) >= 1) {
    
        
    $buffer = 'points(';
        $tmparg0 = [];
        foreach ($input->points as $arg1) {
        $tmptmppointsarg1 = [];
        foreach ($arg1 as $arg1Value) {
        switch (true) {
            case is_string($arg1Value):
                $disjunctionarg1Value =\var_export($arg1Value, true);
                $tmparg1arg1Value = $disjunctionarg1Value;
                break;
            case is_int($arg1Value):
                $disjunctionarg1Value =\var_export($arg1Value, true);
                $tmparg1arg1Value = $disjunctionarg1Value;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmptmppointsarg1[] = $tmparg1arg1Value;
        }
        $tmppointsarg1 = "[" . implode(", \n", $tmptmppointsarg1) . "]";
        $tmparg0[] = $tmppointsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dropPercent !== null) {
    
        
    $buffer = 'dropPercent(';
        $arg0 =\var_export($input->dropPercent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

