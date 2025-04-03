import m from 'mithril';
import { routes } from './routes/index';

m.route(document.getElementById('app'), '/', routes);
