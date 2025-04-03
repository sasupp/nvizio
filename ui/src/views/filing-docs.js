import m from 'mithril';
import CompanyLayout from '../layouts/company';
import FinancialDocs from '../model/financial-docs';
import Global from '../model/global';

const FinancialDocsView = () => {

    const sectionMapping = {
        'QuarterlyAndHalfYearlyFinancialResultFormatIndAS': 'Income Statement',
        'DisclosureOfGeneralInformationAboutCompany': 'General Information',
        'StatementOfAssetsAndLiabilitiesIndAS': 'Balance Sheet',
        'CashFlowStatementIndirect': 'Cashflow (Indirect)',
        'CashFlowStatementDirect': 'Cashflow (Direct)',
        'QuarterlyAndHalfYearlyFinancialResultFormatBanking': 'Income Statement',
        'StatementOfAssetsAndLiabilitiesBanking': 'Balance Sheet',
        'QuarterlyAndHalfYearlyFinancialResultFormatNBFC': 'Income Statement',
        'StatementOfAssetsAndLiabilitiesNBFC': 'Balance Sheet'
    };
    
    const friendlySectionName = (report_title) => {
        return sectionMapping[report_title] || report_title
    };

    const loadFinancialDocs = (vnode) => {
        Global.menu = 'financials'
        FinancialDocs.load(vnode.attrs.ticker, vnode.attrs.filing_id);
    };

    return {
        oninit: function (vnode) {
            loadFinancialDocs(vnode)
        },

        onbeforeupdate: function (vnode, old) {
            // Check if the ticker has changed
            if (vnode.attrs.ticker !== old.attrs.ticker) {
                loadFinancialDocs(vnode);
            }
            return true; // Allow the component to update
        },

        view: (vnode) => {
            const headers = ['Section'];
            return m(CompanyLayout, vnode.attrs,
                FinancialDocs.list.length === 0 ? null :
                    m('table',
                        m('thead',
                            m('tr',
                                headers.map(header => m('th', header))
                            )
                        ),
                        m('tbody',
                            FinancialDocs.list.map(row =>
                                m('tr',
                                    row.type === 'statement' ? [
                                    m('td', m('a', {
                                        href: `#!/nse/company/${vnode.attrs.ticker}/financials/${row.filing_id}/docs/${row.doc_id}`
                                    }, friendlySectionName(row.title))),
                                    ] : null
                                )
                            )
                        )
                    )
            );
        }
    };
};

export default FinancialDocsView;
