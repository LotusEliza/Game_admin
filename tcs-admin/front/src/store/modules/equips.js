import router from "../../router";

export default {
    state: {
        equips: [],
    },
    getters: {
        items: state => state.equips,
    },
    mutations: {
        SET_EQUIPS(state, equips){
            state.equips = equips;
        },
        DELETE_EQUIP(state, equipID){
            let equips = state.equips.filter(e => e.ID != equipID)
            state.equips = equips;
        },
        UPDATE_EQUIP(state, equip){
            let e = state.equips.find(e => e.ID == equip.ID)
            e = equip;
        }
    },
    actions: {
        async addEquip({commit}, equip) {
            let response = await window.axios.post('/equipments/add/', {
                Title: equip.title,
                Type: equip.type,
                SubType: equip.subType,
                BuyPrice: equip.buyPrice,
                SellPrice: equip.sellPrice,
                Reputation: equip.reputation,
                Damage: equip.damage,
                Armor: equip.armor,
                Air: equip.air,
                Mine: equip.mine,
                Time: equip.time,
                SocketType: equip.socketType,
                Sockets: equip.sockets,
            });
            if (response.status == 200 || response.status == 204) {
                console.log('yes ADDED!');
                await router.push({path: '/equipments'});
            }
        },
        async deleteEquip({commit}, equipID) {
            let response = await window.axios.post('/equipments/remove', {
                ID: equipID,
            });
            if (response.status == 200 || response.status == 204) {
                console.log('yes removed!');
                commit('DELETE_EQUIP', equipID);
            }
        },
        async getEquips({commit}){
            let response =  await window.axios.get('/equipments');
            // debugger
            commit('SET_EQUIPS', response.data.Items);
        },
        async updateEquip({commit}, equip){
            let response =  await window.axios.post('/equipment/update', {
                ID: equip.ID,
                Title: equip.Title,
                Type: equip.Type,
                SubType: equip.SubType,
                BuyPrice: equip.BuyPrice,
                SellPrice: equip.SellPrice,
                Reputation: equip.Reputation,
                Damage: equip.Damage,
                Armor: equip.Armor,
                Air: equip.Air,
                Mine: equip.Mine,
                Time: equip.Time,
                SocketType: equip.SocketType,
                Sockets: equip.Sockets,
            });
            if(response.status == 200 || response.status == 204){
                commit('UPDATE_EQUIP', equip);
                await router.push({path: '/equipments'});
            }
        }
    }
}