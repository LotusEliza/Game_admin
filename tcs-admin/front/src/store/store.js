import Vue from 'vue'
import Vuex from 'vuex'
// import router from '../router'
import actions from './actions'
import norms from './modules/norms'
import equips from './modules/equips'

Vue.use(Vuex)

export const store = new Vuex.Store({
    modules:{
        norms,
        equips,
    },
    state: {
        wallet: [],
        player: [],
    },
    mutations: {
        // SET_PLAYERID(state, playerId){
        //     state.playerId = playerId;
        // },
        SET_PLAYER(state, player){
            console.log('player is set' + player);
            state.player= player;
        }
    },
    actions,
    getters : {
        player: state => state.player,
    }
})