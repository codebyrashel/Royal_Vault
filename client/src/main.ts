import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import App from './App.vue';
import './style.css';
import LandingPage from './pages/LandingPage.vue';
import DashboardPage from './pages/DashboardPage.vue';

const routes = [
  { path: '/', component: LandingPage },
  { path: '/dashboard', component: DashboardPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp(App).use(router).mount('#app');