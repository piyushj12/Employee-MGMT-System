import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import VueToast from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'
import VueCookies from 'vue-cookies';
import Vuelidate from 'vuelidate'
import moment from 'moment';


Vue.use(VueCookies);
Vue.use(Vuelidate)
Vue.use(VueToast);
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
Vue.config.productionTip = false;

Vue.filter('formatDate', function(value) {
  if (value) {
      return  moment(new Date(value)).format('MM/DD/YYYY hh:mm A')
      //return moment(String(value)).format('MM/DD/YYYY hh:mm A')
  }
});
Vue.mixin({
  methods: {
    getTitleCase: str => {
      const titleCase = str
      .toLowerCase()
      .split(' ')
      .map(word => {
        return word.charAt(0).toUpperCase() + word.slice(1);
      })
      .join(' ');
  
    return titleCase;
    }
  }
})
new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
