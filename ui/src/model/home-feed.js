import m from "mithril"

var HomeFeed = {
    list: [],
    errorMessage: "",
    loadSampleFeed: function () {
        HomeFeed.list = [
            {
                ticker: 'KSCL',
                company_name: 'Kaveri Seeds',
                subject: "Subject",
                publish_time: "01 Aug",
                details: "Details",
                attachment: "url"
            }
        ]
    },
    load: function () {
        return m.request({
            method: "GET",
            url: `/nse/feed`
        }).then(function (result) {
            HomeFeed.list = result
            HomeFeed.errorMessage = ""
        }).catch(function(error) {
            HomeFeed.list = []
            HomeFeed.errorMessage = ""
            if (error.response && error.response.message) {
                HomeFeed.errorMessage = error.response.message
            } else {
                HomeFeed.errorMessage = "Error downloading feed from the server."
            }
        })
    }
}

export default HomeFeed;