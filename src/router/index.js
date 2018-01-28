import Vue from 'vue';
import Router from 'vue-router';
import Home from '../pages/Home';
import Projects from '../pages/Projects';
import Contact from '../pages/Contact';

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Home',
            component: Home,
        },
        {
            path: '/projects',
            name: 'Projects',
            component: Projects,
        },
        {
            path: '/contact',
            name: 'Contact',
            component: Contact,
        }
    ]
});