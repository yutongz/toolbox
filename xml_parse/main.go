package main

import (
	"encoding/xml"
	"fmt"

	"github.com/jstemmer/go-junit-report/formatter"
)


var XML = []byte(`
<testsuites>
    <testsuite tests="8" failures="0" time="313.356" name="istio.io/istio/tests/e2e/tests/simple">
        <properties>
            <property name="go.version" value="go1.10.1"/>
        </properties>
        <testcase classname="simple" name="TestSimpleIngress" time="15.650"/>
        <testcase classname="simple" name="TestSvc2Svc" time="21.330"/>
        <testcase classname="simple" name="TestAuth" time="0.730"/>
        <testcase classname="simple" name="TestAuthWithHeaders" time="0.600"/>
        <testcase classname="simple" name="Test503sDuringChanges" time="0.000">
            <skipped message="a_simple_1_test.go:249: https://github.com/istio/istio/issues/1038"/>
        </testcase>
        <testcase classname="simple" name="Test503sWithBadClusters" time="0.000">
            <skipped message="a_simple_1_test.go:294: https://github.com/istio/istio/issues/1038"/>
        </testcase>
        <testcase classname="simple" name="TestNginx" time="44.450"/>
        <testcase classname="simple" name="TestNginx/Reachable" time="16.960"/>
    </testsuite>
</testsuites>`)

func main() {
	var jtss formatter.JUnitTestSuites

	err := xml.Unmarshal(XML, &jtss)
	if err != nil {
		fmt.Println(err)
	}


	fmt.Println(len(jtss.Suites))

	for _, jts := range jtss.Suites {
		fmt.Println(jts.Name)
		for _, jtc := range jts.TestCases {
			fmt.Println(jtc.Name)
		}
	}
}