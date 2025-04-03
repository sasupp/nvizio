import m from 'mithril';
import CompanyLayout from '../layouts/company';
import Shareholding from '../model/shareholding';
import Global from '../model/global';
import { companyUrl } from '../lib/util';

const ShareholdingView = () => {

    const loadShareholdingData = (vnode) => {
        Global.menu = 'shp'
        Shareholding.load(vnode.attrs.ticker);
    };

    return {
        oninit: function (vnode) {
            loadShareholdingData(vnode)
        },

        onbeforeupdate: function (vnode, old) {
            // Check if the ticker has changed
            if (vnode.attrs.ticker !== old.attrs.ticker) {
                loadShareholdingData(vnode);
            }
            return true; // Allow the component to update
        },

        onupdate: function () {
            window.scrollTo(0, 0);
        },

        view: (vnode) => {
            let displaySearchResults = Shareholding.searchResults.length > 0;
            let hasError = Shareholding.errorMessage !== "";
            const headers = ['Date', 'Promoter', 'Public', 'Employee Trusts'];
            return m(CompanyLayout, vnode.attrs,
                hasError ? m("div", m("p", Shareholding.errorMessage)) :
                    displaySearchResults ?
                        m("div",
                            m("ul",
                                Shareholding.searchResults.map(result => {
                                    return m("li", m("a", { href: `#!${companyUrl(result.ticker, "shp")}` }, result.company_name))
                                }
                                ))
                        ) :
                        m('table.container',
                            m('thead',
                                m('tr',
                                    headers.map(header => m('th', header))
                                )
                            ),
                            m('tbody',
                                Shareholding.list.map(row =>
                                    m('tr',
                                        m('td', row.as_on_date),
                                        m('td', row.promoter_group),
                                        m('td', row.public),
                                        m('td', row.employee_trusts)
                                    )
                                )
                            )
                        )
            );
        }
    };
};

export default ShareholdingView;
