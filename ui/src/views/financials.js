import m from 'mithril';
import CompanyLayout from '../layouts/company';
import Financials from '../model/financials';
import Global from '../model/global';
import { companyUrl } from '../lib/util';

const FinancialsView = () => {

    const loadFinancials = (vnode) => {
        Global.menu = 'financials'
        Financials.load(vnode.attrs.ticker);
    };

    return {
        oninit: function (vnode) {
            loadFinancials(vnode)
        },

        onbeforeupdate: function (vnode, old) {
            // Check if the ticker has changed
            if (vnode.attrs.ticker !== old.attrs.ticker) {
                loadFinancials(vnode);
            }
            return true; // Allow the component to update
        },

        onupdate: function () {
            window.scrollTo(0, 0);
        },

        view: (vnode) => {
            let displaySearchResults = Financials.searchResults.length > 0;
            let hasError = Financials.errorMessage !== "";
            const headers = ['FiscalYear', 'FiscalPeriod', 'Consolidated', 'Audited', 'Link'];
            return m(CompanyLayout, vnode.attrs,
                hasError ? m("div", m("p", Financials.errorMessage)) :
                    displaySearchResults ?
                        m("div",
                            m("ul",
                                Financials.searchResults.map(result => {
                                    return m("li", m("a", { href: `#!${companyUrl(result.ticker, "financials")}` }, result.company_name))
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
                                Financials.list.map(row =>
                                    m('tr',
                                        m('td', row.fiscal_year),
                                        m('td', row.fiscal_period),
                                        m('td', row.props.consolidated),
                                        m('td', row.props.audited),
                                        m('td', m('a', {
                                            href: `#!/nse/company/${vnode.attrs.ticker}/financials/${row.filing_id}/docs`
                                        }, 'Link')),
                                    )
                                )
                            )
                        )
            );
        }
    };
};

export default FinancialsView;
