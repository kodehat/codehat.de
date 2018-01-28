import Vue from 'vue';
import App from './App';
import router from './router';
import Buefy from 'buefy';
import 'buefy/lib/buefy.css';

Vue.use(Buefy);

Vue.config.productionTip = false;

new Vue({
    el: '#app',
    router,
    render: h => h(App),
    template: '<App/>',
    components: { App },
});