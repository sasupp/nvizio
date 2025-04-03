import m from 'mithril'

var Financials = {
    list: [],
    searchResults: [],
    errorMessage: "",
    loadSampleFinancials: function (ticker) {
        Financials.list = [
            {
                filing_date: ticker,
                company_name: ticker,
                fiscal_year: 2023,
                fiscal_period: 'Q3',
                props: {
                    consolidated: 'Standalone',
                    audited: 'Un-audited',
                },
                doc: 'link'
            }
        ]
    },
    load: function (ticker) {
        return m.request({
            method: "GET",
            url: `/nse/company/${ticker}/financials`
        }).then(function (result) {
            Financials.list = result
            Financials.searchResults = []
            Financials.errorMessage = ""
        }).catch(function(error) {
            if (error.code === 300) {
                Financials.list = []
                Financials.searchResults = error.response
                Financials.errorMessage = ""
            } else {
                Financials.list = []
                Financials.searchResults = []
                
                if (error.response && error.response.message) {
                    Financials.errorMessage = error.response.message
                } else {
                    Financials.errorMessage = "No results found."
                }
            }
        })
    }
}

export default Financials;