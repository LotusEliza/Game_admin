<template>
    <div class="p-3">
        <div v-if="equip[0]">
            <div class="tile is-ancestor">
                <div class="tile is-2 is-vertical is-parent">
                </div>
                <div class="tile is-vertical is-parent">
                    <!--//********************************NAME*******************************************************-->
                    <div class="tile is-child box">
                        <p class="title">Update Equipment #{{$route.params.tg}}</p>
                        <section>
                            <div class="columns is-mobile is-multiline">
                                <div class="column is-narrow">
                                    <b-field label="Title">
                                        <b-input :value="equip[0].Title"
                                                 @input="updateTitle"
                                                 placeholder="Enter title"
                                                 rounded
                                                 size="is-small"
                                        ></b-input>
                                    </b-field>
                                </div>
                                <div class="column is-narrow">
                                    <b-field label="Type">
                                        <b-select placeholder="Select a type"
                                                  :value="equip[0].Type"
                                                  @input="updateType"
                                                  rounded
                                                  size="is-small"
                                        >
                                            <option v-for="(option, key) in options.type"
                                                    :key="key"
                                            >
                                                {{option}}
                                            </option>
                                        </b-select>
                                    </b-field>
                                </div>
                                <div class="column is-narrow">
                                    <b-field label="Sub Type">
                                        <b-select placeholder="Select a socket type"
                                                  :value="equip[0].SubType"
                                                  @input="updateSubType"
                                                  rounded
                                                  size="is-small"
                                        >
                                            <option v-for="(option, key) in options.subType"
                                                    :key="key"
                                                    :value="option.text"
                                            >
                                                {{option.text}}
                                            </option>
                                        </b-select>
                                    </b-field>
                                </div>
                                <div class="column is-narrow">
                                    <b-field label="Socket Type">
                                        <b-select placeholder="Select a socket type"
                                                  :value="equip[0].SocketType"
                                                  @input="updateSocketType"
                                                  rounded
                                                  size="is-small"
                                        >
                                            <option v-for="(option, key) in options.socketType"
                                                    :key="key"
                                                    :value="option.value"
                                            >
                                                {{option.text}}
                                            </option>
                                        </b-select>
                                    </b-field>
                                </div>
                            </div>
                            <div class="columns is-mobile is-multiline">
                                <div class="column">
                                    <b-field label="Sockets">
                                        <b-numberinput :value="equip[0].Sockets"
                                                       @input="updateSockets"
                                                       placeholder="Enter sockets"
                                                       rounded
                                                       size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                                <div class="column">
                                    <b-field label="Buy price">
                                        <b-numberinput :value="equip[0].BuyPrice"
                                                       @input="updateBuyPrice"
                                                       placeholder="Enter buy price"
                                                       rounded
                                                       size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                                <div class="column">
                                    <b-field label="Sell price">
                                        <b-numberinput
                                                :value="equip[0].SellPrice"
                                                @input="updateSellPrice"
                                                placeholder="Enter sell price"
                                                rounded
                                                size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                            </div>
                            <div class="columns">
                                <div class="column is-3">
                                    <b-field label="Reputation">
                                        <b-numberinput :value="equip[0].Reputation"
                                                 @input="updateReputation"
                                                 placeholder="Enter reputation"
                                                 rounded
                                                 size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                                <div class="column is-3">
                                    <b-field label="Damage">
                                        <b-numberinput :value="equip[0].Damage"
                                                       @input="updateDamage"
                                                       placeholder="Enter damage"
                                                       rounded
                                                       size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                                <div class="column is-3">
                                    <b-field label="Armor">
                                        <b-numberinput :value="equip[0].Armor"
                                                       @input="updateArmor"
                                                       placeholder="Enter armor"
                                                       rounded
                                                       size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                            </div>
                            <div class="columns is-mobile is-multiline">
                                <div class="column">
                                    <b-field label="Air">
                                        <b-numberinput :value="equip[0].Air"
                                                       @input="updateAir"
                                                       placeholder="Enter air"
                                                       rounded
                                                       size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                                <div class="column">
                                    <b-field label="Mine">
                                        <b-numberinput :value="equip[0].Mine"
                                                       @input="updateMine"
                                                       placeholder="Enter mine"
                                                       rounded
                                                       size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                                <div class="column">
                                    <b-field label="Time">
                                        <b-numberinput :value="equip[0].Time"
                                                       @input="updateTime"
                                                       placeholder="Enter time"
                                                       rounded
                                                       size="is-small"
                                        ></b-numberinput>
                                    </b-field>
                                </div>
                            </div>
                        </section>
                        <section>
                            <div class="buttons pt-2" >
                                <b-button  @click="updateEquip()" type="is-success">Submit</b-button>
                                <b-button tag="router-link"
                                          to="/equipments"
                                          type="is-link is-light"
                                >
                                    Back
                                </b-button>
                            </div>
                        </section>
                    </div>
<!--                    <div class="tile is-child">-->
<!--                        <notification :message="message"></notification>-->
<!--                    </div>-->
                </div>
                <div class="tile is-2 is-vertical is-parent">
                </div>
            </div>
        </div>
        <div v-else>{{redPage()}}</div>
    </div>
</template>

<script>

    export default {
        name: 'updateEquip',
        components: {
        },
        data() {
            return {
                options: {
                    "type": ['weapon', 'armor', 'balloon', 'console'],
                    "subType": [
                        {text: 'boer', value: 0},
                        {text: 'nullt', value: 1},
                    ],
                    "socketType": [
                        {text: 'damage', value: 1},
                        {text: 'armor', value: 1},
                        {text: 'hp', value: 2},
                        {text: 'mine', value: 3}],
                },
                // message: '',
            }
        },
        mounted(){
            this.$store.commit('equips/SET_EQUIP_ID', this.$route.params.tg)
        },
        computed: {
            equip() {
                return this.$store.getters['equips/equip'](this.$route.params.tg)
            }
        },
        methods: {
            async updateEquip(){
                try {
                    await   this.$store.dispatch('equips/updateEquip');
                } catch (error) {
                    console.log('error')
                } finally {
                    // await this.$router.push('/equipments');
                }
            },
            updateTitle (e) {
                this.$store.commit('equips/UPDATE_TITLE', e)
            },
            updateType(e){
                this.$store.commit('equips/UPDATE_TYPE', e);
            },
            updateSubType(e){
                this.$store.commit('equips/UPDATE_SUBTYPE', e)
            },
            updateBuyPrice(e){
                this.$store.commit('equips/UPDATE_BUY_PRICE', e)
            },
            updateSellPrice(e){
                this.$store.commit('equips/UPDATE_SELL_PRICE', e)
            },
            updateReputation(e){
                this.$store.commit('equips/UPDATE_REPUTATION', e)
            },
            updateDamage(e){
                this.$store.commit('equips/UPDATE_DAMAGE', e)
            },
            updateArmor(e){
                this.$store.commit('equips/UPDATE_ARMOR', e)
            },
            updateAir(e){
                this.$store.commit('equips/UPDATE_AIR', e)
            },
            updateMine(e){
                this.$store.commit('equips/UPDATE_MINE', e)
            },
            updateTime(e){
                this.$store.commit('equips/UPDATE_TIME', e)
            },
            updateSocketType(e){
                this.$store.commit('equips/UPDATE_SOCKET_TYPE', e)
            },
            updateSockets(e){
                this.$store.commit('equips/UPDATE_SOCKETS', e)
            },
            redPage(){
                console.log('No data!')
                this.$router.push({path: '/equipments'});
            },
        },

    }
</script>

<style>
</style>
