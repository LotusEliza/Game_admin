import { ToastProgrammatic as Toast } from 'buefy'
import Vue from 'vue'

// MUTATION(state, payload) {
//     Vue.set(state.arr, index, payload);
// }

export default {
    namespaced: true,
    state: {
        norms: {},
        playerId: null,
    },
    getters: {
        norms: state => state.norms,
        newNorm: state => state.newNorm,
        tg: state => state.playerId,
    },
    mutations: {
        SET_NORMS(state, norms){
            state.norms = norms;
        },
        DELETE_NORM(state, normDate){
            let norms = state.norms.filter(n => n.Date != normDate);
            state.norms = norms;
        },
        UPDATE_NORM(state, res){
            console.log(state.norms)
            Vue.set(state.norms, res.index, res.norm);
        },
        ADD_NORM(state, norm){
            if(!state.norms){
                state.norms = [norm];
            }else {
                let n = state.norms.concat(norm);
                state.norms = n;
            }
        },
        SET_PLAYERID(state, playerId){
            state.playerId = playerId;
        },
        UPDATE_AMOUNT(state, obj){
            let n = state.norms.find(n => n.Date == obj.Date);
            n.Amount = obj.Amount;
        },
        UPDATE_RESOURCE(state, obj){
            let n = state.norms.find(n => n.Date == obj.Date);
            n.Resource = obj.Resource;
        },
    },
    actions:{
        async getNorms({commit, state}, normTg){
            state.norms = [];
            let response =  await window.axios.get('/player/norms?tg='+normTg);
            commit('SET_NORMS', response.data.Item);
        },
        async getNorm({commit}, obj){
            let response =  await window.axios.get('/player/norm?tg='+obj.norm.Tg+'&date='+obj.norm.Date);
            let res = {'norm': response.data.Item[0], 'index': obj.index};
            commit('UPDATE_NORM', res);
        },
        async deleteNorm({commit, state}, date){
            let norm = state.norms.find(n => n.Date == date);
            let response =  await window.axios.post('/player/norms/remove', {
                Tg: norm.Tg,
                Resource: norm.Resource,
                Date: norm.Date
            });
            console.log(norm.Date)
            if(response.status == 200 || response.status == 204){
                Toast.open({
                    message: 'The norm is removed!',
                    type: 'is-danger'
                });
                commit('DELETE_NORM', date);
            }
        },
        async updateNorm({state}, date){
            let norm = state.norms.find(n => n.Date == date);
            let response =  await window.axios.post('/player/norms/update', {
                Tg: norm.Tg,
                Amount: Number(norm.Amount),
                Resource: norm.Resource,
                Date: norm.Date
            });
            if(response.status == 200 || response.status == 204){
                console.log('updated vuex!')
                Toast.open({
                    message: 'The norm is updated!',
                    type: 'is-success'
                });
            }
        },
        async addNorm({commit}, norm){
            let response =  await window.axios.post('/player/norms/add', {
                Tg: norm.Tg,
                Amount: Number(norm.Amount),
                Resource: norm.Resource,
                Date: norm.Date
            });
            if(response.status == 200 || response.status == 204){
                commit('ADD_NORM', norm);
                Toast.open({
                    message: 'The norm is added!',
                    type: 'is-success'
                })
            }
        },
        setPlayerId({commit}, playerId){
            commit('SET_PLAYERID', playerId);
        },
    }
}
