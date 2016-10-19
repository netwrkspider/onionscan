package steps

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/s-rah/onionscan/report"
)

type JsonReportWriter struct {
        reportFile string
}

func (jrw *JsonReportWriter) Init(outputFile string) {
        jrw.reportFile = outputFile
}

func (jrw *JsonReportWriter) Do(r *report.OnionScanReport) {
	jsonOut, err := r.Serialize()

	if err != nil {
		log.Fatalf("Could not serialize json report %v", err)
	}

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%s\n", jsonOut))

        reportFile := r.HiddenService + "." + jrw.reportFile

	if len(reportFile) > 0 {
		f, err := os.Create(reportFile)

		for err != nil {
			log.Printf("Cannot create report file: %s...trying again in 5 seconds...", err)
			time.Sleep(time.Second * 5)
			f, err = os.Create(reportFile)
		}

		defer f.Close()

		f.WriteString(buffer.String())
	} else {
		fmt.Print(buffer.String())
	}
}
