import {ToastProgrammatic as Toast} from "buefy";
import {callFunction} from "../../utils/filters.js"

export default {
    namespaced: true,
    state: {
        players: [],
        player: [],
        total_players: null,
        isActive: false,
        wallet: [],
        newWallet: [],
        inventory: [],
    },
    getters: {
        players: state => state.players,
        total_players: state => state.total_players,
        player: state => state.player,
        // playerOne(state) {
        //     return keyword => state.players.filter(player =>{
        //         return player.Tg == keyword
        //     });
        // },
        wallet: state => state.wallet,
        inventory: state => state.inventory,

    } ,
    mutations: {
        SET_PLAYERS(state, players){
            state.players = players;
        },
        SET_TOTAL_PLAYERS(state, total_players){
            state.total_players = total_players;
        },
        SET_WALLET(state, wallet){
            state.wallet = wallet;
        },
        ADD_INVENTORY(state, inv){
            if(!state.inventory){
                state.inventory = [inv];
            }else {
                let i = state.inventory.concat(inv)
                state.inventory = i;
            }
        },
        SET_INVENTORY(state, inventory){
            state.inventory = inventory;
        },
        DELETE_INVENTORY(state, key){
            state.inventory.splice(key, 1);
        },
        SET_PLAYER(state, player){
            state.player = player;
        },
    },
    actions: {
//*******************************PLAYER*********************************************
        async getPlayersByPage({commit}, page){
            let response =  await window.axios.get('/players?page='+page);
            // Registered
            response.data.Players.forEach(p => {
                p.Registered = callFunction(p.Registered);
                p.LastActive = callFunction(p.LastActive);
            });
            commit('SET_PLAYERS', response.data.Players);
            commit('SET_TOTAL_PLAYERS', response.data.TotalPlayers);
        },
        async updateName({ state }){
            let response =  await window.axios.post('/player/name',{
                Tg: state.player.Tg,
                Name: state.player.Name,
            });
            if(response.status == 200 || response.status == 204){
                Toast.open({
                    message: 'Name is updated!',
                    type: 'is-success'
                });
            }
        },
//*******************************WALLET*********************************************
        async getWallet({commit}, tg){
            let response =  await window.axios.get('/player/wallet?id='+tg);
            commit('SET_WALLET', response.data);
        },
        async updateWallet({state}){
            let response =  await window.axios.post('/player/wallet', {
                Tg: state.player.Tg,
                Credits: Number(state.wallet.Credits),
                Gold: Number(state.wallet.Gold),
            });
            if(response.status == 200 || response.status == 204){
                Toast.open({
                    message: 'Wallet is updated!',
                    type: 'is-success'
                });
            }
        },
        async updateLocation({state}){
            let response =  await window.axios.post('/player/location/update', {
                Tg: state.player.Tg,
                Location: Number(state.player.Location),
                PosX: Number(state.player.PosX),
                PosY: Number(state.player.PosY),
            });
            if(response.status == 200 || response.status == 204){
                Toast.open({
                    message: 'Location is updated!',
                    type: 'is-success'
                });
            }
        },
        async updateOxigen({state}){
            let response =  await window.axios.post('/player/air/update', {
                Tg: state.player.Tg,
                Air: Number(state.player.Air),
            });
            if(response.status == 200 || response.status == 204){
                Toast.open({
                    message: 'Oxigen is updated!',
                    type: 'is-success'
                })
            }
        },
//*******************************INVENTORY*********************************************
        async addInventory({commit}, inventory){
            let response =  await window.axios.post('/player/inventory/add', {
                tg: inventory.Tg,
                itemtype: inventory.ItemType,
                itemvalue: Number(inventory.ItemValue),
            });
            if(response.status == 200 || response.status == 204){
                // console.log('VUEX')
                // console.log(inventory)
                commit('ADD_INVENTORY', inventory);
            }
        },
        async getInventory({commit}, tg){
            let response =  await window.axios.get('/player/inventory?tg='+tg);
            commit('SET_INVENTORY', response.data.Items);
        },
        async deleteInventory({commit, state}, key){
            console.log(state.inventory[key].ItemType)
            let response =  await window.axios.post('/player/inventory/remove', {
                Tg: state.player.Tg,
                ItemType: state.inventory[key].ItemType,
                ItemValue: Number(state.inventory[key].ItemValue),
            });
            if(response.status == 200 || response.status == 204){
                commit('DELETE_INVENTORY', key);
                Toast.open({
                    message: 'Inventory is removed!',
                    type: 'is-danger'
                })
            }
        },
        async getPlayer({commit}, tg){
            console.log('player get vuex')
            let response =  await window.axios.get('/player?id='+tg);
            // /* eslint-disable no-debugger */
            // debugger
            // /* eslint-enable no-debugger */

            if(response.status == 200 || response.status == 204){
                commit('SET_PLAYER', response.data);
            }
            console.log(response.data)

        },
        async updateInventory({state}, key){
            console.log('Type' + state.inventory[key].ItemType)
            console.log('Val' + state.inventory[key].ItemValue)
            console.log('Tg' + state.player.Tg)

            let response =  await window.axios.post('/player/inventory/update', {
                Tg: state.player.Tg,
                ItemType: state.inventory[key].ItemType,
                ItemValue: state.inventory[key].ItemValue,
            });
            if(response.status == 200 || response.status == 204){
                console.log('updated invent vuex!')
                Toast.open({
                    message: 'Inventory is updated!',
                    type: 'is-success'
                })
            }
        },
        // async updateInventory(key){
        //     /player/inventory/update
        //     console.log('update inventory'+key)
        // },
    }
}
