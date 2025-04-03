import m from 'mithril'

var FinancialDocs = {
    list: [],
    loadSampleFinancialDocs: function (ticker, filing_id) {
        FinancialDocs.list = []
    },
    load: function (ticker, filing_id) {
        return m.request({
            method: "GET",
            url: `/nse/company/${ticker}/financials/${filing_id}/docs`
        }).then(function (result) {
            FinancialDocs.list = result
        }).catch(function(error) {
            FinancialDocs.list = []
        })
    }
}

export default FinancialDocs;