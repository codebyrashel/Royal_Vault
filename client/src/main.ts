import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import App from './App.vue';
import './style.css';
import LandingPage from './pages/LandingPage.vue';
import DashboardPage from './pages/DashboardPage.vue';
import CryptoTestPage from './pages/CryptoTestPage.vue';

const routes = [
  { path: '/', component: LandingPage },
  { path: '/dashboard', component: DashboardPage },
  { path: '/crypto-test', component: CryptoTestPage }, // dev route
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp(App).use(router).mount('#app');