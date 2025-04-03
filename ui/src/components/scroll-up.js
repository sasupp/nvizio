import m from 'mithril';

const ScrollUpComponent = () => {
    return {
        view: function () {
            return m('a', {
                id: 'back-to-top',
                href: '#',
                onclick: function (e) {
                    e.preventDefault(); // Prevent default anchor behavior
                    document.documentElement.scrollIntoView({ behavior: 'smooth' }); // Scroll to top smoothly
                }
            }, m('i', { class: 'fas fa-angle-up' }));
        }
    }
};

export default ScrollUpComponent;
