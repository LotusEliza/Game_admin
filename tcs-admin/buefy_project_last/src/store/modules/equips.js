// import router from "../../routes";
import {ToastProgrammatic as Toast} from "buefy";

export default {
    namespaced: true,
    state: {
        equips: [],
        isActive: false,
        equipId: null,
    },
    getters: {
        items: state => state.equips,
        // equip: state => state.equips.filter(e => e.ID == equipID),
        equip(state) {
            return keyword => state.equips.filter(equip =>{
                return equip.ID == keyword
            });
        }
    },
    mutations: {
        ADD_EQUIP(){
            console.log('add')
        },
        SET_EQUIPS(state, equips){
            state.equips = equips;
        },
        DELETE_EQUIP(state, equipID){
            let equips = state.equips.filter(e => e.ID != equipID)
            state.equips = equips;
        },
        // UPDATE_EQUIP(state, equip){
        //     // let e = state.equips.find(e => e.ID == equip.ID)
        //     // equip = e;
        // },
//*********************************FORM UPDATE EQUIP*****************************
        UPDATE_TITLE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Title = equip;
        },
        UPDATE_TYPE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Type = equip;
        },
        UPDATE_SUBTYPE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.SubType = equip;
        },
        UPDATE_BUY_PRICE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.BuyPrice = equip;
        },
        UPDATE_SELL_PRICE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.SellPrice = equip;
        },
        UPDATE_REPUTATION(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Reputation = equip;
        },
        UPDATE_DAMAGE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Damage = equip;
        },
        UPDATE_ARMOR(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Armor = equip;
        },
        UPDATE_AIR(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Air = equip;
        },
        UPDATE_MINE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Mine = equip;
        },
        UPDATE_TIME(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Time = equip;
        },
        UPDATE_SOCKET_TYPE(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.SocketType = equip;
        },
        UPDATE_SOCKETS(state, equip){
            let e = state.equips.find(e => e.ID ==  state.equipId);
            e.Sockets = equip;
        },

        SET_EQUIP_ID(state, id){
            state.equipId = id;
        }
    },
    actions: {
        async addEquip({commit}, equip) {
            console.log(equip);
            console.log(equip.Title);

            let response = await window.axios.post('/equipments/add', {
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
            if (response.status == 200 || response.status == 204) {
                commit('ADD_EQUIP');
                Toast.open({
                    message: 'The equipment is added!',
                    type: 'is-success'
                });
            }
        },
        async deleteEquip({commit}, equipID) {
            let response = await window.axios.post('/equipments/remove', {
                ID: equipID,
            });
            if (response.status == 200 || response.status == 204) {
                Toast.open({
                    message: 'The equipment is removed!',
                    type: 'is-danger'
                });
                console.log('yes removed!');
                commit('DELETE_EQUIP', equipID);
            }
        },
        async getEquips({commit}){
            let response =  await window.axios.get('/equipments');
            // debugger
            commit('SET_EQUIPS', response.data.Items);
        },
        async updateEquip({ state, getters}){
            let equip =  getters.equip(state.equipId);
            let response =  await window.axios.post('/equipment/update', {
                ID: equip[0].ID,
                Title: equip[0].Title,
                Type: equip[0].Type,
                SubType: equip[0].SubType,
                BuyPrice: equip[0].BuyPrice,
                SellPrice: equip[0].SellPrice,
                Reputation: equip[0].Reputation,
                Damage: equip[0].Damage,
                Armor: equip[0].Armor,
                Air: equip[0].Air,
                Mine: equip[0].Mine,
                Time: equip[0].Time,
                SocketType: equip[0].SocketType,
                Sockets: equip[0].Sockets,
            });
            if(response.status == 200 || response.status == 204){
                console.log('updated')
                Toast.open({
                    message: 'Equipment is updated!',
                    type: 'is-success'
                })
            }
        }
    }
}
