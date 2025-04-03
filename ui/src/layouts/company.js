import m from 'mithril';
import TitleComponent from '../components/title'
import NavComponent from '../components/nav';
import ScrollUpComponent from '../components/scroll-up';

const CompanyLayout = () => {
    return {
        view: function (vnode) {
            return m(".container",
                m(TitleComponent, { text: "Nvizio", ticker: vnode.attrs.ticker }),
                m(NavComponent, { ticker: vnode.attrs.ticker }),
                m(".content", vnode.children),
                m(ScrollUpComponent)
            );
        }
    }
};

export default CompanyLayout;
