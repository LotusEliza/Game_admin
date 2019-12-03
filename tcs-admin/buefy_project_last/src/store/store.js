import Vue from 'vue'
import Vuex from 'vuex'
// import router from '../router'
import players from './modules/players'
import equips from './modules/equips'
import norms from './modules/norms'

Vue.use(Vuex)

export const store = new Vuex.Store({
    strict: true,
    modules:{
        players,
        equips,
        norms,
    },
    state: {

    },
    mutations: {

    },
    getters : {

    }
})
