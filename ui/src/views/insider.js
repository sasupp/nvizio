import m from 'mithril';
import CompanyLayout from '../layouts/company';
import Insider from '../model/insider';
import Global from '../model/global';
import { to_lakhs } from '../lib/util';
import { companyUrl } from '../lib/util';

const InsiderView = () => {

    const loadInsiderData = (vnode) => {
        Global.menu = 'insider'
        Insider.load(vnode.attrs.ticker);
    };

    return {
        oninit: function (vnode) {
            loadInsiderData(vnode)
        },

        onbeforeupdate: function (vnode, old) {
            // Check if the ticker has changed
            if (vnode.attrs.ticker !== old.attrs.ticker) {
                loadInsiderData(vnode);
            }
            return true; // Allow the component to update
        },

        onupdate: function () {
            window.scrollTo(0, 0);
        },

        view: (vnode) => {
            let displaySearchResults = Insider.searchResults.length > 0;
            let hasError = Insider.errorMessage !== "";
            const headers = ['Date', 'Type', 'Value (Lakhs)', 'Person', 'Person Category'];
            return m(CompanyLayout, vnode.attrs,
                hasError ? m("div", m("p", Insider.errorMessage)) :
                    displaySearchResults ?
                        m("div",
                            m("ul",
                                Insider.searchResults.map(result => {
                                    return m("li", m("a", { href: `#!${companyUrl(result.ticker, "insider")}` }, result.company_name))
                                }
                                ))
                        ) :
                        m('table',
                            m('thead',
                                m('tr',
                                    headers.map(header => m('th', header))
                                )
                            ),
                            m('tbody',
                                Insider.list.map(row =>
                                    m('tr',
                                        m('td', row.filing_date),
                                        m('td', { class: row.transaction_type.toUpperCase() === 'SELL' ? 'sell-type' : row.transaction_type.toUpperCase() === 'BUY' ? 'buy-type' : '' }, row.transaction_type),
                                        m('td', to_lakhs(row.transaction_value, 2)),
                                        m('td', row.name),
                                        m('td', row.person_category)
                                    )
                                )
                            )
                        )
            );
        }
    };
};

export default InsiderView;
