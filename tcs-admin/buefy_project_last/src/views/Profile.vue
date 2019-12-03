<template>
    <div class="p-3">
        <b-button tag="router-link"
                  :to="{ path: `/players/profile/${player.Tg}/norms` }"
                  type="is-link is-info"
                  class="m-5"
        >
            Norms
        </b-button>
        <b-button @click="$router.go(-1)" type="is-info" class="m-5">
            Back
        </b-button>
        <div class="columns" v-if="player.Tg">
            <div class="column is-one-fifth">
                        <!--//********************************NAME*******************************************************-->
                        <div class="tile is-child box">
                            <p class="title">Name</p>
                            <div class="content">
                                <b-field>
                                    <b-input  v-model="player.Name"  size="is-small" rounded></b-input>
                                </b-field>
                                <div class="buttons">
                                    <b-button @click="updateName()" type="is-primary"  size="is-small" expanded>Update</b-button>
                                </div>
                            </div>
                            <hr>
                            <!--//********************************WALLET*******************************************************-->
                            <p class="title">Wallet</p>
                            <div class="columns">
                                <div class="column">
                                    <p class="subtitle">Gold</p>
                                    <b-field>
                                        <b-input v-model="wallet.Gold"  size="is-small" rounded></b-input>
                                    </b-field>
                                </div>
                                <div class="column">
                                    <p class="subtitle">Credits</p>
                                    <b-field>
                                        <b-input v-model="wallet.Credits"  size="is-small" rounded>22</b-input>
                                    </b-field>
                                </div>
                            </div>
                            <b-button @click="updateWallet()" type="is-primary"  size="is-small" expanded>Update</b-button>
                        </div>
            </div>
            <div class="column is-one-quarter">
                <div class="tile is-child box">
                    <!--//********************************LOCATION*******************************************************-->
                    <p class="title">Location</p>
                    <div class="columns">
                        <div class="column">
                            <p class="subtitle">X coord</p>
                            <b-field>
                                <b-input v-model="player.PosX"  size="is-small" rounded></b-input>
                            </b-field>
                        </div>
                        <div class="column">
                            <p class="subtitle">Y coord</p>
                            <b-field>
                                <b-input v-model="player.PosY"  size="is-small" rounded></b-input>
                            </b-field>
                        </div>
                    </div>
                    <b-field>
                        <b-select rounded expanded v-model="player.Location"  size="is-small">
                            <option v-for="(option, key) in locations"
                                    :key="key"
                                    :value="key"
                            >
                                {{option}}
                            </option>
                        </b-select>
                    </b-field>
                    <div class="buttons">
                        <b-button @click="updateLocation()"  size="is-small" type="is-primary" expanded>Update</b-button>
                    </div>
                    <hr>
                    <!--//********************************CONSOLE*******************************************************-->
                    <p class="title">Console</p>
                    <b-field>
                        <b-input :value="player.EquipC"  size="is-small" rounded></b-input>
                    </b-field>
                </div>
            </div>
            <div class="column">
                    <div class="tile is-child box">
                        <p class="title">Oxigen</p>
                        <b-field>
                            <b-numberinput v-model="player.Air"  size="is-small" rounded></b-numberinput>
                        </b-field>
                        <div class="columns is-mobile is-multiline">
                            <div class="column">
                                <b-button @click="fillUp()"  size="is-small" type="is-primary" expanded>Fill up</b-button>
                            </div>
                            <div class="column">
                                <b-button @click="updateOxigen()"  size="is-small" type="is-primary" expanded>Update</b-button>
                            </div>
                        </div>
                        <hr>
                        <!--//********************************INVENTORY *******************************************************-->
                        <p class="title">Inventory</p>
                        <b-field>
                            <b-select rounded expanded v-model="selected.ItemType"  size="is-small">
                                <option v-for="(option, key) in firstInputOptions"
                                        :key="key"
                                        :value="option"
                                >
                                    {{option}}
                                </option>
                            </b-select>
                        </b-field>
                        <b-field>
                            <b-select rounded expanded v-model="selected.ItemValue"  size="is-small">
                                <option v-for="(option, key) in secondInputOptions"
                                        :key="key"
                                        :value="key"
                                >
                                    {{option.text}}
                                </option>
                            </b-select>
                        </b-field>
                        <div class="buttons">
                            <b-button @click="addInventory()"  size="is-small" type="is-primary" expanded>Add</b-button>
                        </div>
                        <div v-if="inventory">
                            <ul id="item-list" v-show="inventory">
                                <li v-for="(item, key) of inventory" :key="key">
                                    <div class="columns is-mobile is-multiline is-centered">
                                        <div class="column is-2">
                                            {{ item.ItemType }}
                                        </div>
                                        <div class="column is-5 "  v-show="visible !== key">
                                            <div v-for="(option, key) in secondInputOptions2(item.ItemType)"
                                                 :key="key"
                                                 :value="key"
                                            >
                                                <p v-if="key==item.ItemValue">
                                                    {{option.text}}
                                                </p>
                                            </div>
                                        </div>
                                        <div class="column is-5" v-show="visible === key">
                                            <b-field>
                                                <b-select rounded expanded v-model="item.ItemValue" size="is-small">
                                                    <option v-for="(option, key) in secondInputOptions2(item.ItemType)"
                                                            :key="key"
                                                            :value="key"
                                                    >
                                                        {{option.text}}
                                                    </option>
                                                </b-select>
                                            </b-field>
                                        </div>
                                        <div class="column is-1 m-3">
                                            <b-button @click="removeInventory(key)"
                                                      type="is-primary"
                                                      size="is-small"
                                            >
                                                X
                                            </b-button>
                                        </div>
                                        <div class="column is-2 m-3" v-show="visible !== key">
                                            <b-button @click="updateInventory(key)"
                                                      type="is-primary"
                                                      size="is-small"

                                            >
                                                Update
                                            </b-button>
                                        </div>
                                        <div class="column is-2 m-3" v-show="visible === key">
                                            <b-button @click="saveInventory(key)"
                                                      type="is-primary"
                                                      size="is-small"
                                            >
                                                Save
                                            </b-button>
                                        </div>
                                    </div>
                                </li>
                            </ul>
                        </div>
                        <div v-else>
                        </div>
                    </div>
            </div>
            <div class="column is-one-fifth">
                <div class="tile is-child box">
                    <p class="title">Balance</p>
                    <b-field>
                        <b-input value="22" rounded  size="is-small"></b-input>
                    </b-field>
                </div>
            </div>
        </div>
        <div class="tile is-ancestor" v-else>
            <div class="tile is-3 is-vertical is-parent">
                <div class="tile is-child box">
                    No user with such ID!
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex';

    export default {
        name: 'profile',
        components: {
        },
        data(){
            return{
                locations: ['Left Side', 'Right Side'],
                playerItem: [],
                message: '',
                selected: {
                    Tg: Number(this.$route.params.tg),
                    ItemType: null,
                    ItemValue: null,
                },
                firstInputOptions: ['weapon', 'armor', 'metal', 'gas', 'junk', 'device'],
                options: {
                    weapon: [
                        { value: 1, text: 'HGG' },
                        { value: 2, text: 'ТТ' },
                        { value: 3, text: 'МП-40'},
                    ],
                    armor: [
                        { value: 4, text: 'СКФ-98' },
                        { value: 5, text: 'СКФ-98а' },
                        { value: 6, text: 'СКФ-07'},
                    ],
                    metal: [
                        { value: 7, text: 'Железо' },
                        { value: 8, text: 'Медь' },
                        { value: 9, text: 'Титан'},
                    ],
                    gas: [
                        { value: 10, text: 'Метан' },
                        { value: 11, text: 'Водород' },
                        { value: 12, text: 'Гелий'},
                    ],
                    junk: [
                        { value: 13, text: 'Уголь' },
                        { value: 14, text: 'Песок' },
                        { value: 15, text: 'Вода'},
                        { value: 16, text: 'Камни'},
                        { value: 17, text: 'Кости'},
                        { value: 18, text: 'Шкура'},
                        { value: 19, text: 'Ветки'},
                    ],
                    device: [
                        { value: 20, text: 'Бур-1' },
                        { value: 21, text: 'Портативный NULL-T' },
                    ],
                },
                display: false,
                visible: false,
            }
        },
        mounted() {
                this.$store.dispatch('players/getWallet', this.$route.params.tg);
                this.$store.dispatch('players/getInventory', this.$route.params.tg);
                this.$store.dispatch('players/getPlayer', this.$route.params.tg)
        },
        computed: {
            ...mapGetters("players", [
                'player',
                'players',
                'wallet',
                'inventory',
            ]),
            secondInputOptions(){
                return this.selected.ItemType === 'weapon' ? this.options.weapon :
                       this.selected.ItemType === 'armor' ? this.options.armor :
                       this.selected.ItemType === 'metal' ? this.options.metal :
                       this.selected.ItemType === 'gas' ? this.options.gas :
                       this.selected.ItemType === 'junk' ? this.options.junk :
                       this.selected.ItemType === 'device' ? this.options.device :
                       null
            },
        },
        methods:{
            updateName(){
                this.$store.dispatch('players/updateName');
            },
            updateWallet(){
                this.$store.dispatch('players/updateWallet');
            },
            updateLocation(){
                this.$store.dispatch('players/updateLocation');
            },
            plusOxigen() {
                this.player.Air++
            },
            subOxigen(){
                this.player.Air--
            },
            fillUp(){
                this.player.Air = 25;
                this.$store.dispatch('players/updateOxigen');
            },
            updateOxigen(){
                this.$store.dispatch('players/updateOxigen');
            },
            async addInventory(){
                let copiedObject = Object.assign({}, this.selected);
                try {
                    await   this.$store.dispatch('players/addInventory', copiedObject);
                } catch (error) {
                    console.log('error')
                } finally {
                    this.selected.ItemType = null;
                    this.selected.ItemValue = null;
                }
            },
            removeInventory(key){
                this.$store.dispatch('players/deleteInventory', key);
            },
            updateInventory(key){
                this.visible = key;
            },
            saveInventory(key){
                this.$store.dispatch('players/updateInventory', key);
                this.visible = false;
            },
            secondInputOptions2(type){
                return type === 'weapon' ? this.options.weapon :
                type === 'armor' ? this.options.armor :
                type === 'metal' ? this.options.metal :
                type === 'gas' ? this.options.gas :
                type === 'junk' ? this.options.junk :
                type === 'device' ? this.options.device :
                null
            }
        }
    }
</script>

<style>
    .container{
        padding: 20px;
    }
    .title{
        color: #745cd2;
        font-weight: normal;
    }
     body {
         font-size: 14px;
     }
    hr {
        margin-top: 1rem;
        margin-bottom: 1rem;
        border: 0;
        border-top: 1px solid rgba(0, 0, 0, 0.09);
    }
    .tile.is-child.box{
    background-color: rgba(248, 229, 190, 0.09);
    }
    .m-5{
        margin: 5px;
    }
    p.title{
        font-size: 14;
    }

    .tile.is-child {
        display: flex;
        flex-direction: column;
    }

    .columns.is-mobile > .column.is-5{
        text-align: center;
    }
</style>
