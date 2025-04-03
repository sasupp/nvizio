import m from 'mithril';
import UnitSwitch from '../components/unit-switch';
import CompanyLayout from '../layouts/company';
import FinancialDoc from '../model/financial-doc';
import Global from '../model/global';
import { to_lakhs, to_crores } from '../lib/util';

const FinancialDocView = () => {
    
    const loadFinancialDoc = (vnode) => {
        Global.menu = 'financials'
        FinancialDoc.load(vnode.attrs.ticker, vnode.attrs.filing_id, vnode.attrs.doc_id);
    };

    return {
        oninit: function (vnode) {
            loadFinancialDoc(vnode)
        },

        onbeforeupdate: function (vnode, old) {
            // Check if the ticker has changed
            if (vnode.attrs.ticker !== old.attrs.ticker) {
                loadFinancialDoc(vnode);
            }
            return true; // Allow the component to update
        },

        view: (vnode) => {
            return m(CompanyLayout, vnode.attrs,
                FinancialDoc.rows.length === 0 ? null :
                    m('div',
                        m(UnitSwitch, {
                            checked: FinancialDoc.unit_function === to_crores,
                            label1: 'Lakhs ',
                            label2: 'Crores',
                            callback: FinancialDoc.toggleUnit,
                        }),
                        m('table',
                            m('thead',
                                m('tr',
                                    FinancialDoc.headers.map(header => m('th', m.trust(header)))
                                )
                            ),
                            m('tbody',
                                FinancialDoc.rows.map(row => m('tr', {
                                    class: row.is_total ? 'total-row' : ''
                                }, row.elements.map(cell => m('td', cell))))
                            )
                        )
                    ),
                    
            );
        }
    };
};

export default FinancialDocView;
