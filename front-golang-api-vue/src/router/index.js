import { createRouter, createWebHistory } from "vue-router";
import HomePage from "../views/HomePage.vue";
import JogarPage from "../views/JogarPage.vue";

const routes = [
  { path: "/", component: HomePage },
  { path: "/jogar", component: JogarPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
