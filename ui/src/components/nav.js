import m from 'mithril';
import { companyUrl } from '../lib/util';
import Global from '../model/global';

const NavComponent = () => {
    return {
        view: function (vnode) {
            return m("div",
                m("nav",
                    m("ul",
                        m("li", m(m.route.Link, {
                            href: companyUrl(vnode.attrs.ticker, "feed"),
                            class: Global.menu === "feed" ? "active" : ""
                        }, m("h4", "Feed"))),
                        m("li", m(m.route.Link, {
                            href: companyUrl(vnode.attrs.ticker, "financials"),
                            class: Global.menu === "financials" ? "active" : ""
                        }, m("h4", "Financials"))),
                        m("li", m(m.route.Link, {
                            href: companyUrl(vnode.attrs.ticker, "shp"),
                            class: Global.menu === "shp" ? "active" : ""
                        }, m("h4", "Shareholding"))),
                        m("li", m(m.route.Link, {
                            href: companyUrl(vnode.attrs.ticker, "insider"),
                            class: Global.menu === "insider" ? "active" : ""
                        }, m("h4", "Insider"))),
                    )),
                m("h3", vnode.attrs.ticker.toUpperCase()))
        }
    }
};

export default NavComponent;
