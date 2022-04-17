import Vue from 'vue'
import App from './App.vue'
import router from '@/router';

//import {captureNewImage} from './api';


Vue.config.productionTip = false


//captureNewImage({});


new Vue({
  render: h => h(App),
  router
}).$mount('#app')
