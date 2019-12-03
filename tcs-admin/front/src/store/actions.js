import router from "../router";

export default{

    // async addEquip({commit}, equip) {
    //     let response = await window.axios.post('/equipments/add/', {
    //         Title: equip.title,
    //         Type: equip.type,
    //         SubType: equip.subType,
    //         BuyPrice: equip.buyPrice,
    //         SellPrice: equip.sellPrice,
    //         Reputation: equip.reputation,
    //         Damage: equip.damage,
    //         Armor: equip.armor,
    //         Air: equip.air,
    //         Mine: equip.mine,
    //         Time: equip.time,
    //         SocketType: equip.socketType,
    //         Sockets: equip.sockets,
    //     });
    //     // console.log(norm)
    //     if (response.status == 200 || response.status == 204) {
    //         console.log('yes ADDED!');
    //         await router.push({path: '/equipments'});
    //         // commit('ADD_EQUIP', equip);
    //     }
    // },
    // async deleteEquip({commit}, equipID){
    //     let response =  await window.axios.post('/equipments/remove', {
    //         ID: equipID,
    //     });
    //     if(response.status == 200 || response.status == 204){
    //         console.log('yes removed!');
    //         await router.push({path: '/equipments'});
    //         // commit('DELETE_NORM', norm.Date);
    //     }

        // console.log("Removed!"+item.ID)
        // window.axios.post('/equipments/remove', {
        //     ID: item.ID,
        // })
        //     .then(function (response) {
        //         // console.log(response);
        //     })
        //     .catch(function (error) {
        //         console.log(error);
        //     });
        // setTimeout( () => this.fetchData(), 100);
    // },

    // async getWallet({commit}, playerId){
    //     let response =  await window.axios.get(`/player/wallet?id=${playerId}`);
    //     commit('SET_WALLET', response.data);
    // },
    // async updateWallet({commit}){
    //     let response =  await window.axios.post('/player/wallet', {
    //         Tg: this.playerId,
    //         Credits: this.wallet.Credits,
    //         Gold: this.wallet.Gold,
    //     });
    //     if(response.status == 200 || response.status == 204){
    //         commit('UPDATE_WALLET', norm);
    //     }
    // },
    // async setPlayer({commit}, player){
    //     let response =  await window.axios.get(`/players?page=${currentPage}`);
    //     commit('SET_PLAYER', response.data);
    // },
}