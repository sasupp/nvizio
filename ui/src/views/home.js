import m from 'mithril';
import HomeLayout from '../layouts/home';
import HomeFeed from '../model/home-feed';
import Global from '../model/global';
import { companyUrl } from '../lib/util';

var HomeView = () => {
    var showCompany = false;

    const loadFeedData = (vnode) => {
        showCompany = vnode.attrs.showCompany;
        Global.menu = 'feed'
        HomeFeed.load();
    };

    return {
        oninit: function (vnode) {
            loadFeedData(vnode)
        },

        view: (vnode) => {
            return m(HomeLayout, vnode.attrs, 
                m("div", HomeFeed.list.map(feed => {
                    return m("hgroup", [
                        showCompany ? m("h4", [
                            m("a", { href: `#!${companyUrl(feed.ticker, "feed")}` }, feed.company_name),
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
            )
        }
    }
}

export default HomeView;
