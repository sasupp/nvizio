import m from 'mithril';
import HomeView from '../views/home';
import FeedView from '../views/feed';
import ShareholdingView from '../views/shareholding';
import InsiderView from '../views/insider';
import FinancialsView from '../views/financials';
import FinancialDocsView from '../views/filing-docs';
import FinancialDocView from '../views/filing-doc';

export const routes = {
    '/': {
        render: () => {
            return m(HomeView, { showCompany: true });
        }
    },
    "/nse/company/:ticker/feed": {
        render: (vnode) => {
            return m(FeedView, { ticker: vnode.attrs.ticker, showCompany: false })
        }
    },
    "/nse/company/:ticker/shp": {
        render: (vnode) => {
            return m(ShareholdingView, { ticker: vnode.attrs.ticker })
        }
    },
    "/nse/company/:ticker/insider": {
        render: (vnode) => {
            return m(InsiderView, { ticker: vnode.attrs.ticker })
        }
    },
    "/nse/company/:ticker/financials": {
        render: (vnode) => {
            return m(FinancialsView, { ticker: vnode.attrs.ticker })
        }
    },
    "/nse/company/:ticker/financials/:filing_id/docs": {
        render: (vnode) => {
            return m(FinancialDocsView, { ticker: vnode.attrs.ticker, filing_id: vnode.attrs.filing_id})
        }
    },
    "/nse/company/:ticker/financials/:filing_id/docs/:doc_id": {
        render: (vnode) => {
            return m(FinancialDocView, { ticker: vnode.attrs.ticker, filing_id: vnode.attrs.filing_id, doc_id: vnode.attrs.doc_id})
        }
    },
};