import m from 'mithril';
import Global from '../model/global';

const TitleComponent = () => {
  var text
  var ticker

  const loadState = (vnode) => {
    text = vnode.attrs.text;
    ticker = vnode.attrs.ticker
  };

  return {
    oninit: function (vnode) {
      loadState(vnode)
    },

    view: (vnode) => {
      return m("div",
        m("h1.center-text-margin", m("a", { href: `#!/`}, text)),
        m("form", {
          onsubmit: (e) => {
            e.preventDefault();
            m.route.set(`/nse/company/${ticker}/${Global.menu}`);
          }
        },
          m("fieldset", { role: "group" },
            m("input", {
              type: "search",
              id: "search-input",
              name: "search",
              placeholder: "Ticker",
              autocomplete: "off",
              value: ticker,
              oninput: (e) => {
                ticker = e.target.value;
              }
            }),
            m("input", { type: "submit", value: "Submit" })
          ),
          m("small.center-text", "search using ticker or company name"),
        )
      );
    }
  };
};

export default TitleComponent;
