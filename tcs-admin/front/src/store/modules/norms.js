export default {
    state: {
        norms: [],
        playerId: null,
    },
    getters: {
        norms: state => state.norms,
        tg: state => state.playerId,
    },
    mutations: {
        SET_NORMS(state, norms){
            state.norms = norms;
        },
        DELETE_NORM(state, normDate){
            let norms = state.norms.filter(n => n.Date != normDate)
            state.norms = norms;
        },
        UPDATE_NORM(state, norm){
            let n = state.norms.find(n => n.Date == norm.Date)
            // n = norm;
            state.norms = n;
        },
        ADD_NORM(state, norm){
            // console.log(state.norms)
            if(!state.norms){
                state.norms = norm;
            }else {
                let n = state.norms.concat(norm)
                state.norms = n;
            }
        },
        SET_PLAYERID(state, playerId){
            state.playerId = playerId;
        }
    },
    actions:{
        async getNorms({commit}, normTg){
            let response =  await window.axios.get('/player/norms?tg='+normTg);
            commit('SET_NORMS', response.data.Item);
        },
        async deleteNorm({commit}, norm){
            let response =  await window.axios.post('/player/norms/remove', {
                Tg: norm.Tg,
                Resource: norm.Resource,
                Date: norm.Date
            });
            console.log(norm.Date)
            if(response.status == 200 || response.status == 204){
                commit('DELETE_NORM', norm.Date);
            }
        },
        async updateNorm({commit}, norm){
            let response =  await window.axios.post('/player/norms/update', {
                Tg: norm.Tg,
                Amount: Number(norm.Amount),
                Resource: norm.Resource,
                Date: norm.Date
            });
            if(response.status == 200 || response.status == 204){
                commit('UPDATE_NORM', norm);
            }
        },
        async addNorm({commit}, norm){
            // console.log(norm)
            let response =  await window.axios.post('/player/norms/add', {
                Tg: norm.Tg,
                Amount: Number(norm.Amount),
                Resource: norm.Resource,
                Date: norm.Date
            });
            if(response.status == 200 || response.status == 204){
                commit('ADD_NORM', norm);
            }
        },
        setPlayerId({commit}, playerId){
            commit('SET_PLAYERID', playerId);
        },
    }
}