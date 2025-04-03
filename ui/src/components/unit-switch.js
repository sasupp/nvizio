import m from 'mithril';

const UnitSwitch = {
    view: (vnode) => {
        const { checked, callback, label1, label2 } = vnode.attrs;

        return m('div',
            m('label',
                label1,
                m('input', {
                    role: 'switch',
                    type: 'checkbox',
                    checked: checked,
                    onclick: (e) => {
                        callback(e.target.checked)
                    }
                }),
                label2,
            ),
        );
    }
};

export default UnitSwitch;
