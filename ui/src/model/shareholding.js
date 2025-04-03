import m from 'mithril'

var Shareholding = {
    list: [],
    searchResults: [],
    errorMessage: "",
    loadSampleShareholding: function (ticker) {
        Shareholding.list = [
            {
                as_on_date: ticker,
                promoter_group: 30,
                public: 50,
                employee_trusts: 20
            }
        ]
    },
    load: function (ticker) {
        return m.request({
            method: "GET",
            url: `/nse/company/${ticker}/shp`
        }).then(function (result) {
            Shareholding.list = result
            Shareholding.searchResults = []
            Shareholding.errorMessage = ""
        }).catch(function(error) {
            if (error.code === 300) {
                Shareholding.list = []
                Shareholding.searchResults = error.response
                Shareholding.errorMessage = ""
            } else {
                Shareholding.list = []
                Shareholding.searchResults = []
                
                if (error.response && error.response.message) {
                    Shareholding.errorMessage = error.response.message
                } else {
                    Shareholding.errorMessage = "No results found."
                }
            }
        })
    }
}

export default Shareholding;