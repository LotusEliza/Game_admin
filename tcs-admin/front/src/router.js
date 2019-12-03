import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Players from "./components/Players";
import Profile from "./components/Profile";
import Equipments from "./components/Equipments";
import UpdateEquip from "./components/UpdateEquip";
import AddEquip from "./components/AddEquip";
import Test from "./components/Test";
import UpdateNorm from "./components/UpdateNorm";
import Norms from "./components/Norms";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: "/players",
      name: 'players',
      component: Players,
      meta: { showNavigation: true },
    },
    { path: '/players/players', redirect: '/players' },
    { path: '/players/profile/players', redirect: '/players' },
    {
      path: '/players/profile',
      name: 'profile',
      component: Profile,
      props: true
    },
    {
      path: '/equipments',
      name: 'equipments',
      component: Equipments,
      props: true
    },
    {
      path: '/equipments/updequip',
      name: 'updequip',
      component: UpdateEquip,
      props: true
    },
    {
      path: '/equipments/addequip',
      name: 'addequip',
      component: AddEquip,
      props: true
    },
    {
      path: '/test',
      name: 'test',
      component: Test,
    },
    {
      path: '/players/profile/norms/update',
      name: 'updatenorm',
      component: UpdateNorm,
    },
      {
          path: '/players/profile/norms',
          name: 'norms',
          component: Norms,
      },
    // {
    //   path: '/players/profile/norms/add',
    //   name: 'addNorm',
    //   component: AddNorm,
    //    props: true
    // },
  ]
})
