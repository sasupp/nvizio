import m from "mithril"

var Feed = {
    list: [],
    searchResults: [],
    errorMessage: "",
    loadSampleFeed: function (ticker) {
        Feed.list = [
            {
                ticker: ticker,
                company_name: ticker,
                subject: "Subject",
                publish_time: "01 Aug",
                details: "Details",
                attachment: "url"
            }
        ]
    },
    load: function (ticker) {
        return m.request({
            method: "GET",
            url: `/nse/company/${ticker}/feed`
        }).then(function (result) {
            Feed.list = result
            Feed.searchResults = []
            Feed.errorMessage = ""
        }).catch(function(error) {
            if (error.code === 300) {
                Feed.list = []
                Feed.searchResults = error.response
                Feed.errorMessage = ""
            } else {
                Feed.list = []
                Feed.searchResults = []
                
                if (error.response && error.response.message) {
                    Feed.errorMessage = error.response.message
                } else {
                    Feed.errorMessage = "No results found."
                }
            }
        })
    }
}

export default Feed;