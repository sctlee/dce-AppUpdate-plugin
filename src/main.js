import Vue from 'vue';
import VueResource from 'vue-resource';
import ElementUI from 'element-ui';
import * as lodash from './helpers/lodash';
import 'element-ui/lib/theme-default/index.css';
import App from './App.vue';

Vue.use(ElementUI);
Vue.use(VueResource);
Vue.use(lodash);

/* eslint-disable no-new */
new Vue({
  el: '#app',
  render: h => h(App),
});
