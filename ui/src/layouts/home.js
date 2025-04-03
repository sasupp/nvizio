import m from 'mithril';
import TitleComponent from '../components/title'
import ScrollUpComponent from '../components/scroll-up';

var HomeLayout = () => {
    return {
        view: function (vnode) {
            return m(".container", [
                m(TitleComponent, { text: "Nvizio", ticker: vnode.attrs.ticker }),
                m(".content", vnode.children),
                m(ScrollUpComponent)
            ]);
        }
    }
}

export default HomeLayout;