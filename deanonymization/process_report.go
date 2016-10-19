package deanonymization

import (
	"github.com/s-rah/onionscan/config"
	"github.com/s-rah/onionscan/report"
	"github.com/s-rah/onionscan/utils"
)

type ExtractIdentifierStep struct {
        osc *config.OnionScanConfig
}

func (eis *ExtractIdentifierStep) Init(osc *config.OnionScanConfig) {
        eis.osc = osc
}

func (eis *ExtractIdentifierStep) Do(osreport *report.OnionScanReport) {
	anonreport := new(report.AnonymityReport)
	ApacheModStatus(osreport, anonreport, eis.osc)
	CheckExposedDirectories(osreport, anonreport, eis.osc)
	PGPContentScan(osreport, anonreport, eis.osc)
	MailtoScan(osreport, anonreport, eis.osc)
	CheckExif(osreport, anonreport, eis.osc)
	PrivateKey(osreport, anonreport, eis.osc)
	ExtractGoogleAnalyticsID(osreport, anonreport, eis.osc)
	ExtractGooglePublisherID(osreport, anonreport, eis.osc)
	ExtractBitcoinAddress(osreport, anonreport, eis.osc)
	GetOnionLinks(osreport, anonreport, eis.osc)
	CommonCorrelations(osreport, anonreport, eis.osc)
	utils.RemoveDuplicates(&anonreport.RelatedOnionServices)
	utils.RemoveDuplicates(&anonreport.RelatedClearnetDomains)
	utils.RemoveDuplicates(&anonreport.IPAddresses)
	utils.RemoveDuplicates(&anonreport.EmailAddresses)
	utils.RemoveDuplicates(&anonreport.AnalyticsIDs)
	utils.RemoveDuplicates(&anonreport.BitcoinAddresses)
	utils.RemoveDuplicates(&anonreport.LinkedOnions)
	osreport.SimpleReport = report.SummarizeToSimpleReport(osreport.HiddenService, anonreport)
	osreport.AnonymityReport = anonreport
}
