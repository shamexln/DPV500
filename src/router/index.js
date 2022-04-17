//引入vue-router路由插件
import VueRouter from "vue-router";
//引入Vue
import Vue from "vue";

Vue.use(VueRouter);

import Home from '@/pages/Home'

export default new VueRouter({
    routers:[
        {
            path:"/Home",
            component:Home
        }
    ]
})