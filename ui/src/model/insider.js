import m from 'mithril'

var Insider = {
    list: [],
    searchResults: [],
    errorMessage: "",
    loadSampleInsider: function (ticker) {
        Insider.list = [
            {
                ticker: ticker,
                filing_date: ticker,
                transaction_type: "Sell",
                transaction_value: 102,
                name: "Suman",
                person_category: "Director"
            }
        ]
    },
    load: function (ticker) {
        return m.request({
            method: "GET",
            url: `/nse/company/${ticker}/insider`
        }).then(function (result) {
            Insider.list = result
            Insider.searchResults = []
            Insider.errorMessage = ""
        }).catch(function(error) {
            if (error.code === 300) {
                Insider.list = []
                Insider.searchResults = error.response
                Insider.errorMessage = ""
            } else {
                Insider.list = []
                Insider.searchResults = []
               
                if (error.response && error.response.message) {
                    Insider.errorMessage = error.response.message
                } else {
                    Insider.errorMessage = "No results found."
                }
            }
        })
    }
}

export default Insider;