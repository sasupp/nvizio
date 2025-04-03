import m from 'mithril';
import CompanyLayout from '../layouts/company';
import Feed from '../model/feed';
import { companyUrl } from '../lib/util';
import Global from '../model/global';

const FeedView = () => {
    var showCompany = false;

    const loadFeedData = (vnode) => {
        showCompany = vnode.attrs.showCompany;
        Global.menu = 'feed'
        Feed.load(vnode.attrs.ticker);
    };

    return {
        oninit: function (vnode) {
            loadFeedData(vnode)
        },

        onbeforeupdate: function (vnode, old) {
            // Check if the ticker has changed
            if (vnode.attrs.ticker !== old.attrs.ticker) {
                loadFeedData(vnode);
            }
            return true; // Allow the component to update
        },

        onupdate: function () {
            window.scrollTo(0, 0);
        },

        view: (vnode) => {
            let displaySearchResults = Feed.searchResults.length > 0;
            let hasError = Feed.errorMessage !== "";
            return m(CompanyLayout, vnode.attrs,
                hasError ? m("div", m("p", Feed.errorMessage)) :
                    displaySearchResults ?
                        m("div",
                            m("ul",
                                Feed.searchResults.map(result => {
                                    return m("li", m("a", { href: `#!${companyUrl(result.ticker, "feed")}` }, result.company_name))
                                }
                                ))
                        ) :
                        m("div", Feed.list.map(feed => {
                            return m("hgroup", [
                                showCompany ? m("h4", [
                                    m("a", { href: companyUrl(feed.ticker, "feed") }, feed.company_name),
                                ]) : null,
                                m("h5", feed.subject),
                                m("small", feed.publish_time),
                                m("blockquote", [
                                    m("span", feed.details),
                                    m("br"),
                                    m("small", [
                                        m("a", { href: feed.attachment, rel: "noopener noreferrer" }, "Link")
                                    ])
                                ])
                            ]);
                        }))
            );
        }
    };
};

export default FeedView;
